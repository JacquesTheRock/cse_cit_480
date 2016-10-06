package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"treview.com/bloom/entity"
	"treview.com/bloom/auth"
	"treview.com/bloom/util"
	"treview.com/bloom/user"
)


func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		postUsers(w, r)
	}
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	users,_ := user.SearchUsers(entity.User{})
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(users)
	
}
func postUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.FormValue("username")
	pass := r.FormValue("password")
	email := r.FormValue("email")
	name := r.FormValue("name")
	location := r.FormValue("location")
	status := "OK"
	u := entity.User{}
	if uid == "" {
		status = "Required field: username"
	} else if pass == "" {
		status = "Required field: password"
	} else if email == "" {
		status = "Required field: email"
	} else if name == "" {
		status = "Required field: name"
	}
	if uid == "" || pass == "" || email == "" || name == "" {
		util.PrintInfo(status)
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(u)
		return
	} else {
		hash,salt,err := auth.CreateHash(pass,"SHA512")
		if err != nil {
			util.PrintError("Create Hash of password Failed")
			util.PrintError(err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		u, err := user.CreateUser(uid,email,name,location,hash,salt)
		if err != nil {
			util.PrintError("Posting User Failed")
			util.PrintError(err)
		}
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(u)
	}
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
	vars := mux.Vars(r)
	uid := vars["uid"]
	user,_ := user.GetUser(entity.User{ID: uid})
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}
func putUsersUid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	decoder := json.NewDecoder(r.Body)
	var u entity.User
	err := decoder.Decode(&u)
	out := entity.User{}
	if err != nil {
		util.PrintError("Bad request body, expected user JSON")
		util.PrintError(err)
	}
	if u.ID == uid {
		orig,_ := user.GetUser(u)
		if u.DisplayName == "" {
			u.DisplayName = orig.DisplayName
		}
		if u.Email == "" {
			u.Email = orig.Email
		}
		if u.Location == "" {
			u.Location = orig.Location
		}
		out,err = user.UpdateUser(u)
	}

	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
