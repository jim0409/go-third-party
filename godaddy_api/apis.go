package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func domains(customerId string, apiKey string, apiSec string) ([]map[string]interface{}, error) {
	edpts := "domains"
	url := fmt.Sprintf("%v/%v", godaddyProdUrl, edpts)
	auth := fmt.Sprintf("sso-key %v:%v", apiKey, apiSec)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Shopper-Id", customerId)
	req.Header.Set("Authorization", auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Not a good request with response code : %v\n", resp.StatusCode)
	}

	// read response body to ioutil; may assume response body as json format with a domain array
	m := make([]map[string]interface{}, 1)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &m); err != nil {
		return nil, err
	}

	return m, nil
}
