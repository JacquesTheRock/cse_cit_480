package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"treview.com/bloom/entity"
	"treview.com/bloom/auth"
	"treview.com/bloom/util"
)


func updateUser(u entity.User) (entity.User, error) {
	const qBase = "UPDATE users SET name = $1, email = $2, location = $3 WHERE id = $4"
	_, err := util.Database.Exec(qBase, u.DisplayName, u.Email, u.Location, u.ID)
	if err != nil {
		util.PrintError("Failure to Update User")
		return entity.User{}, err
	}
	u, err = getUser(u)
	if err != nil {
		util.PrintError("Failure to find user data")
		return entity.User{}, err
	}
	return u, nil
	
}

func createUser(uid string, email string, name string, location string, hash []byte, salt []byte) (entity.User, error) {
	const qBase = "INSERT INTO users(id,email,name,location,hash,salt,algorithm) VALUES ($1,$2,$3,$4,$5,$6,'SHA512')"
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

func searchUsers(u entity.User) ([]entity.User, error) {
	const qBase = "SELECT id,email,name,location FROM users"
	queryVars := make([]interface{},0)
	out := make([]entity.User,0)
	query := " WHERE "
	endQuery := qBase
	if u.ID != "" {
		queryVars = append(queryVars,u.ID)
		query = query + "id LIKE $" + string(len(queryVars)) + " "
	} else {
		if u.DisplayName != "" {
			queryVars = append(queryVars,u.DisplayName)
			query = query + "name LIKE $" + string(len(queryVars)) + " "
		}
		if len(queryVars) > 0 {
			endQuery = qBase + query
		}
	}
	rows, err := util.Database.Query(endQuery, queryVars...)
	defer rows.Close()
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID,&e.Email,&e.DisplayName,&e.Location)
		if err != nil {
			util.PrintError("Unable to read user")
			util.PrintError(err)
		}
		out = append(out,e)
	}

	return out,nil
}

func getUser(u entity.User) (entity.User, error) {
	const qBase = "SELECT id,email,name,location FROM users WHERE id = $1"
	out := entity.User{}
	rows, err := util.Database.Query(qBase, u.ID)
	defer rows.Close()
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID,&e.Email,&e.DisplayName,&e.Location)
		if err != nil {
			util.PrintError("Unable to read user")
			util.PrintError(err)
		}
		out = e
	}

	return out,nil
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
	users,_ := searchUsers(entity.User{})
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(users)
	
}
func postUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("username")
	pass := r.FormValue("password")
	email := r.FormValue("email")
	name := r.FormValue("name")
	location := r.FormValue("location")
	status := "OK"
	u := entity.User{}
	if user == "" {
		status = "Required field: username"
	} else if pass == "" {
		status = "Required field: password"
	} else if email == "" {
		status = "Required field: email"
	} else if name == "" {
		status = "Required field: name"
	}
	if user == "" || pass == "" || email == "" || name == "" {
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
		u, err := createUser(user,email,name,location,hash,salt)
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
	user,_ := getUser(entity.User{ID: uid})
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
		orig,_ := getUser(u)
		if u.DisplayName == "" {
			u.DisplayName = orig.DisplayName
		}
		if u.Email == "" {
			u.Email = orig.Email
		}
		if u.Location == "" {
			u.Location = orig.Location
		}
		out,err = updateUser(u)
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
