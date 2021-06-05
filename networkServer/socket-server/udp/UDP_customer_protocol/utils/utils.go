package utils

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}
