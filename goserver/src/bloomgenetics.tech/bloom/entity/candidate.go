package entity

type Candidate struct {
	ID        int64   `json:"id"`
	CrossID   int64   `json:"crossId"`
	ProjectID int64   `json:"projectId"`
	ImageID   int64   `json:"imageId"`
	Note      string  `json:"note"`
	Traits    []Trait `json:"traits"`
}
