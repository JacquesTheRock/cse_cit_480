package entity

type Candidate struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"projectId"`
	Name      string `json:"name"`
	Parent1ID int64  `json:"parent1"`
	Parent2ID int64  `json:"parent2"`
	Traits	[]Trait	`json:"traits"`
}
