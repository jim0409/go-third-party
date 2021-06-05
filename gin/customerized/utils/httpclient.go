package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"

	. "go-third-party/gin/customerized/logger"
)

// 使用 Singleton 避免client做copy
func RetriveHttpClient() *Client {
	once.Do(func() {
		httpClient = httpClient
	})
	return httpClient
}

var (
	httpClient *Client
	once       sync.Once
)

type Client struct {
	HTTPClient     *http.Client
	RetryAfterMins retryAfterMins
	RetryMax       int

	CheckForRetry checkForRetry
}

func init() {
	httpClient = newClient()
}

const (
	MaxIdleConns        int = 60000
	MaxIdleConnsPerHost int = 10000
	IdleConnTimeout     int = 150
	RequestTimeout      int = 30
	RetryMaxTimes       int = 3
	InitRetryWaitMins   int = 1
)

/*
	有需要帶上特殊header或是特殊處理的httpClient可以定義在這邊，供後面直接做調度
*/
// Customerize Get method
func (c *Client) Get(url string) (*http.Response, error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Key", "secretKey")
	return c.Do(req)
}

// Customerize Post method
func (c *Client) Post(url, bodyType string, body io.ReadSeeker) (*http.Response, error) {
	req, err := NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	req.Header.Set("Key", "secretKey")
	return c.Do(req)
}

type retryAfterMins func(int) int

func defaultDetryAfterMins(i int) int {
	return 2*i + InitRetryWaitMins
}

type checkForRetry func(resp *http.Response, err error) (bool, error)

func defaultRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}

	if resp.StatusCode == 0 || resp.StatusCode >= 500 {
		return true, nil
	}

	return false, nil
}

func newClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(RequestTimeout) * time.Second,
					KeepAlive: time.Duration(RequestTimeout) * time.Second,
				}).DialContext,

				MaxIdleConns:        MaxIdleConns,
				MaxIdleConnsPerHost: MaxIdleConnsPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Duration(RequestTimeout) * time.Second,
		},
		RetryMax:       RetryMaxTimes,
		RetryAfterMins: defaultDetryAfterMins,
		CheckForRetry:  defaultRetryPolicy,
	}
}

type request struct {
	body io.ReadSeeker
	*http.Request
}

func NewRequest(method, url string, body io.ReadSeeker) (*request, error) {
	var rcBody io.ReadCloser
	if body != nil {
		rcBody = ioutil.NopCloser(body)
	}

	httpReq, err := http.NewRequest(method, url, rcBody)
	if err != nil {
		return nil, err
	}

	return &request{body, httpReq}, nil
}

func (c *Client) Do(req *request) (*http.Response, error) {
	count := 0
	retryAfterMins := 0
	for {
		if req.Body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v\n", err)
			}
		}

		resp, err := c.HTTPClient.Do(req.Request)

		checkOk, checkErr := c.CheckForRetry(resp, err)
		if err != nil {
			Log.Error(fmt.Sprintf("Err %s %s request failed: %v\n", req.Method, req.URL, err))
		}

		if !checkOk {
			if checkErr != nil {
				err = checkErr
			}
			return resp, err
		}

		// going to retry, consume any response to reuse the connection
		if err == nil {
			c.drainBody(resp.Body)
		}

		if count != c.RetryMax {
			retryAfterMins = defaultDetryAfterMins(count)
			count++
		} else {
			break
		}
		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		Log.Info(fmt.Sprintf("DEBUG %s: retrying in %d mins(s) %d left)", desc, retryAfterMins, count))
		time.Sleep(time.Duration(retryAfterMins) * time.Minute)

	}
	return nil, fmt.Errorf("%s %s giving up after %d attemps", req.Method, req.URL, count)
}

// Try to read the response body so we can reuse this connection
func (c *Client) drainBody(body io.ReadCloser) {
	defer body.Close()

	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, 10))
	if err != nil {
		Log.Error(fmt.Sprintf("Error reading response body: %v\n", err))
	}
}
