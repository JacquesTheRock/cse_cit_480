package entity

type User struct {
	ID          string  `json:"id"`
	Email string `json:"email"`
	DisplayName string `json:"name"`
	Location string `json:"location"`
	Growzone string `json:"growzone"`
}
