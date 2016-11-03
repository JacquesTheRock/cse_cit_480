package auth

import (
	"bloomgenetics.tech/bloom/util"
	"encoding/base64"
)

func LogoutUser(auth string) error {
	uid, token := ParseAuthorization(auth)
	const qBase = "DELETE FROM logins WHERE user_id = $1 and key = $2"
	b, err1 := base64.URLEncoding.DecodeString(token)
	if err1 != nil {
		util.PrintDebug(err1)
		util.PrintError("Base64 conversion to Bytea failed")
		return err1
	}
	_, err := util.Database.Exec(qBase, uid, b)
	if err != nil {
		util.PrintDebug(err)
		return err
	}
	return nil
}
