package entity

type ApiData struct {
	Code   int64       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
