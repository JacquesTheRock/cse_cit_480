package auth

import (
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"errors"
	"strings"
	"treview.com/bloom/entity"
	"treview.com/bloom/util"
)

func VerifyPermissions(auth string) bool {
	uid, _ := ParseAuthorization(auth)
	u := CheckAuth(auth)
	return u.ID == uid
}

func searchToken(uid string, token string) (entity.UserLogin, error) {
	output := entity.UserLogin{}
	const qBase = "SELECT user_id,name,key FROM logins WHERE user_id = $1 AND key = $2"
	b, err1 := base64.URLEncoding.DecodeString(token)
	if err1 != nil {
		util.PrintError(err1)
		util.PrintError("Base64 conversion to Bytea failed")
		return output, err1
	}
	rows, err := util.Database.Query(qBase, uid, b)
	if err != nil {
		util.PrintError(err)
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
			util.PrintError(err)
			util.PrintInfo("Failure to Find Matching Token")
		}
	}
	return output, nil
}

func ParseAuthorization(authLine string) (string, string) {
	parts := strings.Split(authLine, " ")
	if len(parts) < 2 {
		util.PrintError("Fail to parse Authorization")
		return "", ""
	}
	auth := parts[1]
	data, err := base64.URLEncoding.DecodeString(auth)
	if err != nil {
		util.PrintError(err)
		util.PrintError("Fail to parse Authorization")
		return "Guest", ""
	}
	s := string(data)
	out := strings.Split(s, ":")
	if len(out) == 2 {
		return out[0], out[1]
	} else {
		util.PrintError(err)
		return "Guest", ""
	}
}

func createToken(uid string) (string, error) {
	const qBase = "INSERT INTO logins(user_id,key) VALUES ($1,$2)"
	b := make([]byte, 32)
	_, err := rand.Read(b) //Make the actual token
	if err != nil {
		util.PrintError(err)
		return "", err
	}
	_, err = util.Database.Exec(qBase, uid, b)
	if err != nil {
		util.PrintError(err)
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)
	return token, nil
}

func LoginUser(user string, pass string) (entity.UserLogin, error) {
	u := entity.UserLogin{
		"",
		"Guest",
		"",
	}
	const qBase = "SELECT id,name,hash,salt,algorithm FROM users WHERE id = $1"
	var id, name, hash, salt, algorithm, checkHash string
	rows, err := util.Database.Query(qBase, user)
	if err != nil {
		util.PrintError(err)
		util.PrintError("User ID not found: " + user)
		return u, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &hash, &salt, &algorithm)
		if err != nil {
			util.PrintError(err)
		}
	}
	if id == "" {
		return u, errors.New("User does not exist: " + user)
	}
	switch algorithm {
	case "SHA512":
		b := sha512.Sum512_256([]byte(pass + salt))
		checkHash = string(b[:])
	case "PLAIN":
		checkHash = pass + salt
	default:
		return u, errors.New("Invalid password algorithm: " + algorithm)
	}
	if strings.Compare(hash, checkHash) == 0 { //Hashes Check out
		u.Token, err = createToken(user)
		u.ID = user
		u.DisplayName = name
		if err != nil {
			util.PrintError("Failure to Create Token for: " + user)
		}
	}
	return u, err
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
		util.PrintError(err)
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

func LogoutUser(auth string) error {
	uid, token := ParseAuthorization(auth)
	const qBase = "DELETE FROM logins WHERE user_id = $1 and key = $2"
	b, err1 := base64.URLEncoding.DecodeString(token)
	if err1 != nil {
		util.PrintError(err1)
		util.PrintError("Base64 conversion to Bytea failed")
		return err1
	}
	_, err := util.Database.Exec(qBase, uid, b)
	if err != nil {
		util.PrintError(err)
		return err
	}
	return nil
}
