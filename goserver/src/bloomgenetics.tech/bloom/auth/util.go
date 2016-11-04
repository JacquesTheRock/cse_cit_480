package auth

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"errors"
	"strings"
)

func searchToken(uid string, token string) (entity.UserLogin, error) {
	output := entity.UserLogin{}
	const qBase = "SELECT user_id,name,key FROM logins WHERE user_id = $1 AND key = $2"
	b, err1 := base64.URLEncoding.DecodeString(token)
	if err1 != nil {
		util.PrintDebug(err1)
		util.PrintError("Base64 conversion to Bytea failed")
		return output, err1
	}
	rows, err := util.Database.Query(qBase, uid, b)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Query Failed")
		return output, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, dname sql.NullString
		var b []byte
		err = rows.Scan(&id, &dname, &b)
		token := base64.URLEncoding.EncodeToString(b)
		if id.Valid {
			output.ID = id.String
		}
		if dname.Valid {
			output.DisplayName = dname.String
		}
		output.Token = token
		if err != nil {
			util.PrintDebug(err)
			util.PrintInfo("Failure to Find Matching Token")
		}
	}
	return output, nil
}

func ParseAuthorization(authLine string) (string, string) {
	parts := strings.Split(authLine, " ")
	if len(parts) < 2 {
		if authLine != "" {
			util.PrintError("Fail to parse Authorization")
		} else {
			util.PrintDebug("Anonymous User")
		}
		return "", ""
	}
	auth := parts[1]
	data, err := base64.URLEncoding.DecodeString(auth)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Fail to parse Authorization")
		return "Guest", ""
	}
	s := string(data)
	out := strings.Split(s, ":")
	if len(out) == 2 {
		return out[0], out[1]
	} else {
		util.PrintDebug(err)
		return "Guest", ""
	}
}

func CreateHash(password string, algorithm string) ([]byte, []byte, error) {
	var salt []byte
	var hash []byte
	if len(password) < 6 {
		out := errors.New("Password too short")
		util.PrintError(out)
		return nil, nil, out
	}
	n, err := rand.Read(salt)
	if err != nil || n != len(salt) {
		util.PrintError("Error getting random Salt")
		util.PrintDebug(err)
		return nil, nil, err
	}
	switch algorithm {
	case "SHA512":
		c := append([]byte(password), salt...)
		h := sha512.Sum512_256(c)
		hash = h[:]
	}
	return hash, salt, nil
}
