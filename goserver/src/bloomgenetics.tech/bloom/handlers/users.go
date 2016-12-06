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
	q := entity.User{}
	qString := r.URL.Query()
	if qString != nil {
		id := qString.Get("id")
		name := qString.Get("name")
		if id != "" {
			q.ID = "%" + id + "%"
		}
		if name != "" {
			q.DisplayName = "%" + name + "%"
		}
	}
	//	q.DisplayName = r.URL.Query().Get("name")
	//	q.Location = r.URL.Query().Get("location")
	//	q.Season = r.URL.Query().Get("season")
	//	q.Growzone = r.URL.Query().Get("growzone")
	//	q.Specialty = r.URL.Query().Get("specialty")
	users, _ := user.SearchUsers(q)
	out.Data = users
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

}
func postUsers(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	nU := entity.User{}
	r.ParseForm()
	nU.ID = r.FormValue("username")
	nU.Email = r.FormValue("email")
	nU.DisplayName = r.FormValue("name")
	nU.Location = r.FormValue("location")
	nU.Growzone = r.FormValue("growzone")
	nU.Season = r.FormValue("season")
	nU.Specialty = r.FormValue("specialty")
	pass := r.FormValue("password")
	status := "OK"
	u := entity.User{}
	if nU.ID == "" {
		out.Status = "Required field: username"
	} else if pass == "" {
		out.Status = "Required field: password"
	} else if nU.Email == "" {
		out.Status = "Required field: email"
	}
	if nU.ID == "" || pass == "" || nU.Email == "" {
		util.PrintInfo(status)
		out.Code = code.MISSINGFIELD
		out.Data = u
	}
	if out.Code == 0 {
		err := nU.Validate()
		if err != nil {
			out.Status = err.Error()
			out.Code = code.INVALIDFIELD
			out.Data = u
		}
	}
	if out.Code == 0 {
		hash, salt, err := auth.CreateHash(pass, "SHA512")
		if err != nil {
			util.PrintError("Create Hash of password Failed")
			util.PrintDebug(err)
			out.Code = code.INVALIDFIELD
			out.Status = "Bad password"
		} else {
			u, err = user.CreateUser(nU, hash, salt)
			if err != nil {
				util.PrintError("Posting User Failed")
				util.PrintDebug(err)
				out.Code = code.UNDEFINED
				out.Status = "Unable to create that user"
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
		util.PrintDebug(err)
		out.Code = code.INVALIDFIELD
		out.Status = "Invalid JSON body in request"
	}
	err = u.Validate()
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = err.Error()
	}
	if out.Code == 0 && u.ID == uid {
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
		q := project.QueryProject{}
		q.ID.Valid = true
		q.ID.Int64 = role.ProjectID
		p, err := project.GetProject(q)
		if err == nil {
			p.Role = auth.GetRole(uid, role.ProjectID).Name
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

func UsersUidProjectsPid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidProjectsPid(w, r)
	case "DELETE":
		deleteUsersUidProjectsPid(w, r)
	}
}

func getUsersUidProjectsPid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Project ID must be numeric"
	}
	if pid == 0 {
		out.Code = code.INVALIDFIELD
		out.Status = "Cannot delete root project"
	}
	p := entity.Project{}
	if out.Code == 0 {
		q := user.QueryProjectRole{UID: uid}
		q.PID.Int64 = pid
		q.PID.Valid = true
		roles, _ := user.SearchProjects(q)
		if len(roles) > 1 {
			out.Code = code.UNDEFINED
			out.Status = "Project ID returned multiple projects, aborting"
		} else if len(roles) == 0 {
			out.Code = code.UNDEFINED
			out.Status = "Project ID not found"
		}
		if out.Code == 0 {
			role := roles[0]
			q := project.QueryProject{}
			q.ID.Valid = true
			q.ID.Int64 = role.ProjectID
			p, err = project.GetProject(q)
			if err != nil {
				out.Status = "Could not get Project Data"
				out.Code = 100
			}
			p.Role = auth.GetRole(uid, p.ID).Name
			out.Data = p
		}
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func deleteUsersUidProjectsPid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Project ID must be numeric"
	}
	if out.Code == 0 {
		q := user.QueryProjectRole{UID: uid}
		q.PID.Int64 = pid
		q.PID.Valid = true
		roles, _ := user.SearchProjects(q)
		if len(roles) > 1 {
			out.Code = code.UNDEFINED
			out.Status = "Project ID returned multiple projects, aborting"
		} else if len(roles) == 0 {
			out.Code = code.UNDEFINED
			out.Status = "Project ID not found"
		}
		if out.Code == 0 {
			role := roles[0]
			err = auth.DeleteRole(role)
			if err != nil {
				out.Status = "Could not remove from project"
				out.Code = 100
			}
		}
	}
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
	sender := auth.GetLogin(r.Header.Get("Authorization"))
	uid := vars["uid"]
	decoder := json.NewDecoder(r.Body)
	var m entity.Mail
	err := decoder.Decode(&m)
	out.Data = entity.Mail{}
	if err != nil {
		util.PrintError("Bad request body, expected mail JSON")
		out.Code = code.INVALIDFIELD
		out.Status = "Not JSON body"
		util.PrintDebug(err)
	}
	if m.Dest == "" {
		m.Dest = uid
	}
	if m.Src == "" {
		m.Src = sender.ID
	}
	if m.Src != sender.ID {
		out.Code = code.INVALIDFIELD
		out.Status = "Username mismatch. Spoofing not supported"
	}
	if out.Code == 0 && m.Dest == uid {
		out.Data, err = user.PostMail(m)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Failed to send mail"
		}
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
				util.PrintDebug(err)
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
				util.PrintDebug(err)
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
