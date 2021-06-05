package api1

type Response struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    bool   `json:"Data"`
}

func GetExternalApi()
