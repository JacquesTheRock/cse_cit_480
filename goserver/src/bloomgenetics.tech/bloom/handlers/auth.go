package handlers

import (
	authlib "bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/code"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"encoding/json"
	"net/http"
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
	out := entity.ApiData{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	auth := r.Header.Get("Authorization")
	u := authlib.CheckAuth(auth)
	out.Data = u
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func postAuth(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	uid := r.FormValue("user")
	pass := r.FormValue("password")
	u := entity.UserLogin{}
	if uid == "" || pass == "" {
		out.Code = code.MISSINGFIELD
		out.Status = "Missing username or password"
	} else {
		var err error
		u, err = authlib.LoginUser(uid, pass)
		if err != nil {
			util.PrintError("Failure to Login User")
			util.PrintError(err)
			out.Status = "Failure to log-in"
		}
	}
	out.Data = u
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteAuth(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	auth := r.Header.Get("Authorization")
	u := authlib.CheckAuth(auth)
	out.Data = u
	if u.ID == "" { //Not logged in as a user
		w.Header().Set("WWW-Authenticate", "Basic realm=\"User\"")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		out.Code = code.INVALIDSTATE
		encoder := json.NewEncoder(w)
		encoder.Encode(out)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authlib.LogoutUser(auth)    //Delete the token
	u = authlib.CheckAuth(auth) //Verify that the token is invalidated
	out.Data = u
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
