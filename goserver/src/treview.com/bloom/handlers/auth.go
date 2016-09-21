package handlers

import (
	//	"fmt"
	"net/http"
	"encoding/json"
	"treview.com/bloom/entity"
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

func loginUser(user string, pass string) entity.UserLogin {
	u := entity.UserLogin{
		-1,
		"Guest",
		"",
	}
	return u
}

func checkToken(token string) entity.UserLogin {
	u := entity.UserLogin{
		-1,
		"Guest",
		"",
	}
	return u
}

func deleteToken(token string) {
}

func getAuth(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	u := checkToken(token)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
func postAuth(w http.ResponseWriter, r *http.Request) {
	u := loginUser("name","pass")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
func deleteAuth(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	u := checkToken(token)
	if u.ID < 0 { //Not logged in as a user
		w.Header().Set("WWW-Authenticate", "Basic realm=\"User\"")
		encoder := json.NewEncoder(w)
		encoder.Encode(u)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	deleteToken(token)//Delete the token
	u = checkToken(token)//Verify that the token is invalidated
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
