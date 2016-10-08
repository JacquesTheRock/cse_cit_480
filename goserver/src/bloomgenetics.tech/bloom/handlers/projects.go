package handlers

import (
	authlib "bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/project"
	"bloomgenetics.tech/bloom/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjects(w, r)
	case "POST":
		postProjects(w, r)
	}
}
func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//TODO: Confirm logged in
	//token := r.Header.Get("Authorization")
	p, _ := project.SearchProjects(entity.Project{})
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}
func postProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	token := r.Header.Get("Authorization")
	ctype := r.Header.Get("Content-type")
	uid, _ := authlib.ParseAuthorization(token)
	p := entity.Project{}
	if authlib.VerifyPermissions(token) {
		e := entity.Project{}
		var err error
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				e = entity.Project{Description: "Invalid JSON Posted"}
				util.PrintError("Unable to decode json")
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
			e.Description = r.FormValue("description")
			e.Visibility, err = strconv.ParseBool(r.FormValue("public"))
			if err != nil {
				util.PrintError(err)
				e.Visibility = false
			}
		}
		if e.Name != "" {
			p, err = project.NewProject(uid, e)
		}
	} else {
		util.PrintInfo("User Access denied")
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)

}

func ProjectsPid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPid(w, r)
	case "PUT":
		putProjectsPid(w, r)
	case "DELETE":
		deleteProjectsPid(w, r)
	}
}
func getProjectsPid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	pArray, _ := project.SearchProjects(entity.Project{ID: pid})
	p := entity.Project{}
	if len(pArray) != 1 {
		p.Description = "Invalid Project Selected"
	} else {
		p = pArray[0]
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}
func putProjectsPid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
	}
	ctype := r.Header.Get("Content-type")
	e := entity.Project{}
	switch ctype {
	case "application/json":
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&e)
		if err != nil {
			e = entity.Project{Description: "Invalid JSON Posted"}
			util.PrintError("Unable to decode json")
			util.PrintError(err)
		}
	default:
		r.ParseForm()
		e.Name = r.FormValue("name")
		e.Description = r.FormValue("description")
		e.Visibility, err = strconv.ParseBool(r.FormValue("public"))
		if err != nil {
			util.PrintError(err)
			e.Visibility = false
		}
	}
	pArray, _ := project.SearchProjects(entity.Project{ID: pid})
	p := entity.Project{}
	if len(pArray) == 1 {
		o := pArray[0]
		if e.Name == "" {
			e.Name = o.Name
		}
		if e.Description == "" {
			e.Description = o.Name
		}
		p, _ = project.UpdateProject(e)
	} else {
		p.Description = "ID mismatch"
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}
func deleteProjectsPid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
	}
	p := entity.Project{ID: pid}
	project.DeleteProject(p)
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidTraits(w, r)
	case "POST":
		postProjectsPidTraits(w, r)
	}
}
func getProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	t := [10]entity.Trait{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
func postProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidTraitsTid(w, r)
	case "POST":
		putProjectsPidTraitsTid(w, r)
	case "DELETE":
		deleteProjectsPidTraitsTid(w, r)
	}
}
func getProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	t := entity.Trait{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
func putProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func ProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrosses(w, r)
	case "POST":
		postProjectsPidCrosses(w, r)
	}
}
func getProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := [10]entity.Cross{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func postProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCid(w, r)
	case "PUT":
		putProjectsPidCrossesCid(w, r)
	case "DELETE":
		deleteProjectsPidCrossesCid(w, r)
	}
}
func getProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := entity.Cross{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func putProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCidCandidates(w, r)
	case "POST":
		postProjectsPidCrossesCidCandidates(w, r)
	}
}
func getProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := [10]entity.Candidate{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func postProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCidCandidatesCnid(w, r)
	case "PUT":
		putProjectsPidCrossesCidCandidatesCnid(w, r)
	case "DELETE":
		deleteProjectsPidCrossesCidCandidatesCnid(w, r)
	}
}
func getProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := entity.Candidate{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func putProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTreview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsTreview(w, r)
	}
}
func getProjectsTreview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTreviewCid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsTreviewCid(w, r)
	}
}
func getProjectsTreviewCid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
