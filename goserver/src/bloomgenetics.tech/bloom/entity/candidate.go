package entity

type Candidate struct {
	ID        int64   `json:"id"`
	CrossID   int64   `json:"crossId"`
	ProjectID int64   `json:"projectId"`
	Traits    []Trait `json:"traits"`
}
