package auth

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func VerifyPermissions(auth string) bool {
	uid, _ := ParseAuthorization(auth)
	u := CheckAuth(auth)
	return u.ID == uid
}
func CheckAuth(auth string) entity.UserLogin {
	u := entity.UserLogin{
		"",
		"Guest",
		"",
	}
	uid, token := ParseAuthorization(auth)
	result, err := searchToken(uid, token)
	if err != nil {
		util.PrintError("Failure to Auth Token for: " + uid)
		return u
	}
	return result
}
