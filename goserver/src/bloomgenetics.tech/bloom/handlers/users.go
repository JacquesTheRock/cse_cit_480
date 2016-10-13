package handlers

import (
	"bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/code"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/project"
	"bloomgenetics.tech/bloom/user"
	"bloomgenetics.tech/bloom/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	out := entity.ApiData{}
	users, _ := user.SearchUsers(entity.User{})
	out.Data = users
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

}
func postUsers(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
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
		out.Status = status
		out.Code = code.MISSINGFIELD
		out.Data = u
	} else {
		hash, salt, err := auth.CreateHash(pass, "SHA512")
		if err != nil {
			util.PrintError("Create Hash of password Failed")
			util.PrintError(err)
			out.Code = code.INVALIDFIELD
		} else {
			u, err = user.CreateUser(uid, email, name, location, hash, salt)
			if err != nil {
				util.PrintError("Posting User Failed")
				util.PrintError(err)
				out.Code = code.UNDEFINED
			}
		}
		out.Data = u
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	user, _ := user.GetUser(entity.User{ID: uid})
	out.Data = user
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func putUsersUid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	decoder := json.NewDecoder(r.Body)
	var u entity.User
	err := decoder.Decode(&u)
	out.Data = entity.User{}
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
		out.Data, err = user.UpdateUser(u)
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
	out := entity.ApiData{}
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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	m, _ := user.GetMails(entity.Mail{Dest: uid})
	out.Data = m
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func postUsersUidMail(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	decoder := json.NewDecoder(r.Body)
	var m entity.Mail
	err := decoder.Decode(&m)
	out.Data = entity.Mail{}
	if err != nil {
		util.PrintError("Bad request body, expected mail JSON")
		util.PrintError(err)
	}
	if m.Dest == "" {
		m.Dest = uid
	}
	if m.Dest == uid {
		out.Data, err = user.PostMail(m)
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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	mid, err := strconv.ParseInt(vars["mid"], 10, 64)
	if err != nil {
		util.PrintError("Invalid Mail ID")
		out.Code = code.INVALIDFIELD
		out.Status = "Invalid Mail ID"
	}
	m, _ := user.GetMailByID(entity.Mail{ID: mid, Dest: uid})
	out.Data = m
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func putUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	mid, err := strconv.ParseInt(vars["mid"], 10, 64)
	if err != nil {
		util.PrintError("Invalid Mail ID")
		out.Code = code.INVALIDFIELD
		out.Status = "Invalid Mail ID"
	} else {
		mArray, _ := user.GetMailByID(entity.Mail{ID: mid, Dest: uid})
		out.Data = entity.Mail{}
		if len(mArray) != 1 {
			util.PrintError("Can't reply to many message")
			out.Status = "Couldn't reply to many message"
			out.Code = code.UNDEFINED
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
				out.Data, err = user.ReplyMail(n)
			}
			if err != nil {
				util.PrintError(err)
				out.Code = code.UNDEFINED
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	mid, err := strconv.ParseInt(vars["mid"], 10, 64)
	if err != nil {
		util.PrintError("Invalid Mail ID")
		out.Code = code.INVALIDFIELD
		out.Status = "Mail ID must be an integer"
	}
	mArray, _ := user.GetMailByID(entity.Mail{ID: mid, Dest: uid})
	out.Data = entity.Mail{}
	if len(mArray) != 1 {
		util.PrintError("Matched multiple Mails with 1 MID in context that should only match 1")
		out.Status = "Couldn't Delete message"
	}
	user.DeleteMail(mArray[0])
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
