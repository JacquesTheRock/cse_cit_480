package entity

type Trait struct {
	ID         int64   `json:"id"`
	Project_ID int64   `json:"project_id"`
	Name       string  `json:"name"`
	Weight     float64 `json:"weight"`
	Type       string  `json:"type"`
	Type_ID    int64   `json:"type_id"`
}
