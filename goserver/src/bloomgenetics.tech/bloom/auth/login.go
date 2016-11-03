package auth

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strings"
)

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
		util.PrintDebug(err)
		util.PrintError("User ID not found: " + user)
		return u, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &hash, &salt, &algorithm)
		if err != nil {
			util.PrintDebug(err)
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

func createToken(uid string) (string, error) {
	const qBase = "INSERT INTO logins(user_id,key) VALUES ($1,$2)"
	b := make([]byte, 32)
	_, err := rand.Read(b) //Make the actual token
	if err != nil {
		util.PrintDebug(err)
		return "", err
	}
	_, err = util.Database.Exec(qBase, uid, b)
	if err != nil {
		util.PrintDebug(err)
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)
	return token, nil
}
