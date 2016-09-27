package handlers

import (
	//"fmt"
	"strings"
	"encoding/base64"
	"crypto/sha512"
	"crypto/rand"
	"net/http"
	"encoding/json"
	"treview.com/bloom/entity"
	"treview.com/bloom/util"
	"database/sql"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAuth(w, r)
	case "POST":
		postAuth(w, r)
	case "DELETE":
		deleteAuth(w, r)
	}
}

func searchToken(uid string, token string) (entity.UserLogin, error) {
	output := entity.UserLogin{}
	const qBase = "SELECT user_id,name,key FROM logins WHERE user_id = $1 AND key = $2"
	b,err1 := base64.StdEncoding.DecodeString(token)
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
		var id,dname sql.NullString
		var b []byte
		err = rows.Scan(&id,&dname,&b)
		token := base64.StdEncoding.EncodeToString(b)
		if id.Valid {
			output.ID = id.String
		}
		if dname.Valid {
			output.DisplayName = dname.String
		}
		output.Token = token
		if err != nil {
			util.PrintError(err)
		}
	}
	return output, nil
}

func parseAuthorization(auth string) (string,string) {
	data, err := base64.URLEncoding.DecodeString(auth)
	if err != nil {
		util.PrintError(err)
		return "Guest",""
	}
	s := string(data)
	out := strings.Split(s,":")
	if len(out) == 2 {
		return out[0],out[1]
	} else {
		util.PrintError(err)
		return "Guest",""
	}
}

func createToken(uid string) (string,error) {
	const qBase = "INSERT INTO logins(user_id,key) VALUES ($1,$2)"
	b := make([]byte, 32)
	_,err := rand.Read(b) //Make the actual token
	if err != nil {
		util.PrintError(err)
		return "", err
	}
	_,err = util.Database.Exec(qBase, uid, b)
	if err != nil {
		util.PrintError(err)
		return "", err
	}
	token := base64.StdEncoding.EncodeToString(b)
	return token,nil
}

func loginUser(user string, pass string) (entity.UserLogin,error) {
	u := entity.UserLogin{
		"",
		"Guest",
		"",
	}
	const qBase = "SELECT name,hash,salt,algorithm FROM users WHERE id = $1"
	var name,hash,salt,algorithm,checkHash string
	rows, err := util.Database.Query(qBase,user)
	if err != nil {
		util.PrintError(err)
		util.PrintError("User ID not found: " + user)
		return u, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name,&hash,&salt,&algorithm)
		if err != nil {
			util.PrintError(err)
		}
	}
	switch algorithm {
		case "SHA512":
			b := sha512.Sum512_256([]byte(pass + salt))
			checkHash = base64.StdEncoding.EncodeToString(b[:])
		case "PLAIN":
			checkHash = pass + salt
	}
	if strings.Compare(hash,checkHash) == 0 { //Hashes Check out
		u.Token, err = createToken(user)
		u.ID = user
		u.DisplayName = name
		if err != nil {
			util.PrintError("Failure to Create Token for: " + user)
		}
	}
	return u, err
}

func checkAuth(auth string) entity.UserLogin {
	u := entity.UserLogin{
		"",
		"Guest",
		"",
	}
	uid,token := parseAuthorization(auth)
	result, err := searchToken(uid,token)
	if err != nil {
		util.PrintError("Failure to Auth Token for: " + uid)
		return u
	}
	return result
}

func logoutUser(auth string) error {
	uid,token := parseAuthorization(auth)
	const qBase = "DELETE FROM logins WHERE user_id = $1 and key = $2"
	b,err1 := base64.StdEncoding.DecodeString(token)
	if err1 != nil {
		util.PrintError(err1)
		util.PrintError("Base64 conversion to Bytea failed")
		return err1
	}
	_,err := util.Database.Exec(qBase, uid, b)
	if err != nil {
		util.PrintError(err)
		return err
	}
	return nil
}

func getAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	auth := r.Header.Get("Authorization")
	u := checkAuth(auth)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
func postAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	u,err := loginUser(r.FormValue("user"),r.FormValue("password"))
	if err != nil {
		util.PrintError("Failure to Login User")
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
func deleteAuth(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	u := checkAuth(auth)
	if u.ID == "" { //Not logged in as a user
		w.Header().Set("WWW-Authenticate", "Basic realm=\"User\"")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(u)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	logoutUser(auth)//Delete the token
	u = checkAuth(auth)//Verify that the token is invalidated
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
