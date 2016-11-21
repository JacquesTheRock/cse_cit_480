package entity

import (
	"bloomgenetics.tech/bloom/util"
	"errors"
	"regexp"
)

type User struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"name"`
	Location    string `json:"location"`
	Growzone    string `json:"growzone"`
	Season      string `json:"season"`
	Specialty   string `json:"specialty"`
	About       string `json:"about"`
}

func (u User) validateID() bool {
	const idRXP = `[a-z]+[a-z0-9@._]+`
	resp, err := regexp.Match(idRXP, []byte(u.ID))
	if err != nil {
		util.PrintDebug(err)
		return false
	}
	return resp
}

func (u User) validateEmail() bool {
	const emRXP = `[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`
	resp, err := regexp.Match(emRXP, []byte(u.Email))
	if err != nil {
		util.PrintDebug(err)
		return false
	}
	return resp
}

func (u User) validateAbout() bool {
	const aboutRXP = `[A-Za-z0-9.%+-]*`
	resp, err := regexp.Match(aboutRXP, []byte(u.About))
	if err != nil {
		util.PrintDebug(err)
		return false
	}
	return resp
}

func (u User) Validate() error {
	if !u.validateID() {
		return errors.New("Invalid username")
	}
	if !u.validateEmail() {
		return errors.New("Invalid email")
	}
	if !u.validateAbout() {
		return errors.New("Invalid About me characters")
	}
	return nil
}
