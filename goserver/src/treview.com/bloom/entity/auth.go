package entity

type UserLogin struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"name"`
	Token string `json:"token"`
}
