package entity

type Image struct {
	ID   int64  `json:"iid"`
	Data string `json:"data"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}
