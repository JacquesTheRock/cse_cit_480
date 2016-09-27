package entity

import (
	//"database/sql"
	//"treview.com/bloom/util"
)

type UserLogin struct {
	ID          string  `json:"id"`
	DisplayName string `json:"name"`
	Token string `json:"token"`
}

