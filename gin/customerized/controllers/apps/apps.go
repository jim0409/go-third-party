package apps

import (
	"fmt"
	"io/ioutil"

	"go-third-party/gin/customerized/utils"
)

func GetResponse(config interface{}) (string, error) {
	/*
		Do some pre-processing before send a request
	*/
	// assert url := http://35.241.127.22
	url := config.(map[string]interface{})["url"].(string)
	return sendRequest(url)
}

func sendRequest(url string) (string, error) {
	resp, err := utils.RetriveHttpClient().Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hi apps responde %v", string(data)), nil
}
