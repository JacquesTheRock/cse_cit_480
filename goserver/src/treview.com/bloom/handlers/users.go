package handlers

import (
	"encoding/json"
	"net/http"
	"treview.com/bloom/entity"
	"treview.com/bloom/auth"
	"treview.com/bloom/util"
)

func createUser(uid string, email string, name string, location string, hash []byte, salt []byte) (entity.User, error) {
	const qBase = "INSERT INTO Users(id,email,name,location,hash,salt,algorithm) VALUES ($1,$2,$3,$4,$5,$6,'SHA512')"
	user := entity.User{}
	_, err := util.Database.Exec(qBase, uid, email,name,location,hash,salt)
	if err != nil {
		util.PrintError("createUser Function")
		util.PrintError(err)
		return user, err
	}
	user.ID = uid
	user.Email = email
	user.DisplayName = name
	user.Location = location
	return user, nil
}

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		postUsers(w, r)
	}
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func postUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("username")
	pass := r.FormValue("password")
	email := r.FormValue("email")
	name := r.FormValue("name")
	location := r.FormValue("location")
	if user == "" {
		w.WriteHeader(http.StatusNotAcceptable)	
		return
	} else if pass == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else if email == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else if name == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	hash,salt,err := auth.CreateHash(pass,"SHA512")
	if err != nil {
		util.PrintError("Create Hash of password Failed")
		util.PrintError(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	u, err := createUser(user,email,name,location,hash,salt)
	if err != nil {
		util.PrintError("Posting User Failed")
		util.PrintError(err)
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}

func UsersUid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUid(w, r)
	case "PUT":
		putUsersUid(w, r)
	}
}
func getUsersUid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func putUsersUid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUidProjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidProjects(w, r)
	}
}
func getUsersUidProjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUidMail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidMail(w, r)
	case "POST":
		postUsersUidMail(w, r)
	}
}
func getUsersUidMail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func postUsersUidMail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidMailMid(w, r)
	case "PUT":
		putUsersUidMailMid(w, r)
	case "DELETE":
		deleteUsersUidMailMid(w, r)
	}
}
func getUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func putUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
