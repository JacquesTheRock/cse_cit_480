package entity

import (
	"bloomgenetics.tech/bloom/util"
	"errors"
	"regexp"
)

type Project struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  bool   `json:"public"`
	Role        string `json:"role"`
}

func (e Project) validateName() bool {
	const nameRXP = `[a-zA-Z]+[a-zA-Z0-9_]+`
	resp, err := regexp.Match(nameRXP, []byte(e.Name))
	if err != nil {
		util.PrintDebug(err.Error())
		return false
	}
	return resp
}

func (e Project) Validate() error {
	if !e.validateName() {
		return errors.New("Invalid Name")
	}
	return nil
}
