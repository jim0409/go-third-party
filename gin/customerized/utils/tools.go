package utils

import "fmt"

// ComposeUrl 用於組合querystring參數 e.g. map[string]interface{}{"t":"v1", "i":123} =>回傳 "test=v1&i=123"
func ComposeUrl(params map[string]interface{}) string {
	if len(params) < 1 {
		return ""
	}

	url := ""

	isFirst := true
	for k, v := range params {
		if v == nil {
			continue
		}

		vStr := fmt.Sprintf("%v", v)
		if vStr == "" || vStr == "null" || vStr == "NULL" || vStr == "nil" {
			continue
		}
		if isFirst != false {
			url += k + "=" + vStr
			isFirst = false
			continue
		}

		url = url + "&" + k + "=" + vStr
	}

	return url
}
