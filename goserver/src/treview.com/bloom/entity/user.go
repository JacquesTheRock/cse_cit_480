package entity

import (
	"encoding/json"
)

type User struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"name"`
}
