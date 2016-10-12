package handlers

import (
	"bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/project"
	"bloomgenetics.tech/bloom/user"
	"bloomgenetics.tech/bloom/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ApiData struct {
	Code   int64
	Status string
	Data   interface{}
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
	users, _ := user.SearchUsers(entity.User{})
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
		hash, salt, err := auth.CreateHash(pass, "SHA512")
		if err != nil {
			util.PrintError("Create Hash of password Failed")
			util.PrintError(err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		u, err := user.CreateUser(uid, email, name, location, hash, salt)
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
	user, _ := user.GetUser(entity.User{ID: uid})
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
		orig, _ := user.GetUser(u)
		if u.DisplayName == "" {
			u.DisplayName = orig.DisplayName
		}
		if u.Email == "" {
			u.Email = orig.Email
		}
		if u.Location == "" {
			u.Location = orig.Location
		}
		out, err = user.UpdateUser(u)
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
	out := ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	roles, _ := user.SearchProjects(user.QueryProjectRole{UID: uid})
	pArray := make([]entity.Project, 0)
	for _, role := range roles {
		p, err := project.GetProject(entity.Project{ID: role.PID})
		if err == nil {
			pArray = append(pArray, p)
		} else {
			out.Status = "Could not find some projects"
			out.Code = 100
		}
	}
	out.Data = pArray
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
	vars := mux.Vars(r)
	uid := vars["uid"]
	m, _ := user.GetMails(entity.Mail{Dest: uid})
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(m)
}
func postUsersUidMail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	decoder := json.NewDecoder(r.Body)
	var m entity.Mail
	err := decoder.Decode(&m)
	out := entity.Mail{}
	if err != nil {
		util.PrintError("Bad request body, expected mail JSON")
		util.PrintError(err)
	}
	if m.Dest == "" {
		m.Dest = uid
	}
	if m.Dest == uid {
		out, err = user.PostMail(m)
	}

	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

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
	vars := mux.Vars(r)
	uid := vars["uid"]
	mid, err := strconv.ParseInt(vars["mid"], 10, 64)
	if err != nil {
		util.PrintError("Invalid Mail ID")
	}
	m, _ := user.GetMailByID(entity.Mail{ID: mid, Dest: uid})
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(m)
}
func putUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	mid, err := strconv.ParseInt(vars["mid"], 10, 64)
	if err != nil {
		util.PrintError("Invalid Mail ID")
	}
	mArray, _ := user.GetMailByID(entity.Mail{ID: mid, Dest: uid})
	out := entity.Mail{}
	if len(mArray) != 1 {
		util.PrintError("Can't reply to many message")
		out = entity.Mail{Message: "Couldn't reply to message"}
	} else {
		m := mArray[0]
		decoder := json.NewDecoder(r.Body)
		var n entity.Mail
		err = decoder.Decode(&n)
		if err != nil {
			util.PrintError("Bad request body, expected mail JSON")
			util.PrintError(err)
		}
		if n.Subject == "" {
			n.Subject = "Re:" + m.Subject
		}
		n.Prev = m.ID
		n.Src = m.Dest
		n.Dest = m.Src
		if n.Src == uid {
			out, err = user.ReplyMail(n)
		}
		if err != nil {
			util.PrintError(err)
		}
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	mid, err := strconv.ParseInt(vars["mid"], 10, 64)
	if err != nil {
		util.PrintError("Invalid Mail ID")
	}
	mArray, _ := user.GetMailByID(entity.Mail{ID: mid, Dest: uid})
	out := entity.Mail{}
	if len(mArray) != 1 {
		util.PrintError("Can't reply to many message")
		out = entity.Mail{Message: "Couldn't Delete message"}
	}
	user.DeleteMail(mArray[0])
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
