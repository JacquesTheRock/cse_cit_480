package handlers

import (
	//"fmt"
	"encoding/json"
	"net/http"
	authlib "treview.com/bloom/auth"
	"treview.com/bloom/util"
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

func getAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	auth := r.Header.Get("Authorization")
	u := authlib.CheckAuth(auth)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
func postAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	u, err := authlib.LoginUser(r.FormValue("user"), r.FormValue("password"))
	if err != nil {
		util.PrintError("Failure to Login User")
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
func deleteAuth(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	u := authlib.CheckAuth(auth)
	if u.ID == "" { //Not logged in as a user
		w.Header().Set("WWW-Authenticate", "Basic realm=\"User\"")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(u)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authlib.LogoutUser(auth)    //Delete the token
	u = authlib.CheckAuth(auth) //Verify that the token is invalidated
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}
