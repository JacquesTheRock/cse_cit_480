package handlers

import (
	authlib "bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/candidate"
	"bloomgenetics.tech/bloom/code"
	"bloomgenetics.tech/bloom/cross"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/project"
	"bloomgenetics.tech/bloom/trait"
	"bloomgenetics.tech/bloom/treview"
	"bloomgenetics.tech/bloom/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
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
	out := entity.ApiData{}
	query := r.URL.Query()
	q := project.QueryProject{}
	if query.Get("pid") != "" {
		var err error
		q.ID.Int64, err = strconv.ParseInt(query.Get("pid"), 10, 64)
		if err == nil {
			q.ID.Valid = true
		} else {
			util.PrintDebug(err)
		}
	}
	if query.Get("name") != "" {
		q.Name.Valid = true
		q.Name.String = "%" + query.Get("name") + "%"
	}
	if query.Get("location") != "" {
		q.Location.Valid = true
		q.Location.String = "%" + query.Get("location") + "%"
	}
	if query.Get("species") != "" {
		q.Species.Valid = true
		q.Species.String = "%" + query.Get("species") + "%"
	}
	if query.Get("type") != "" {
		q.Type.Valid = true
		q.Type.String = "%" + query.Get("type") + "%"
	}

	p, _ := project.SearchProjects(q)

	out.Data = p
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func postProjects(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	token := r.Header.Get("Authorization")
	ctype := r.Header.Get("Content-type")
	uid, _ := authlib.ParseAuthorization(token)
	out.Data = entity.Project{}
	if authlib.VerifyPermissions(token) {
		e := entity.Project{}
		var err error
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				e = entity.Project{Description: "Invalid JSON Posted"}
				out.Code = code.INVALIDFIELD
				out.Status = "Unable to decode json"
				util.PrintError("Unable to decode json")
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
			e.Description = r.FormValue("description")
			e.Location = r.FormValue("location")
			e.Species = r.FormValue("species")
			e.Type = r.FormValue("type")
			e.Visibility, err = strconv.ParseBool(r.FormValue("public"))
			if err != nil {
				util.PrintDebug(err)
				e.Visibility = false
			}
		}
		if e.Name == "" {
			out.Code = code.MISSINGFIELD
			out.Status = "Projects must have a name"
		}
		if out.Code == 0 {
			err = e.Validate()
			if err != nil {
				out.Code = code.INVALIDDATA
				out.Status = "Project name contains invalid characters"
			}
		}
		if out.Code == 0 {
			out.Data, err = project.NewProject(uid, e)
		}
	} else {
		util.PrintInfo("User Access denied")
		out.Code = code.ACCESSDENIED
		out.Status = "You don't have permission"
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

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
	out := entity.ApiData{}
	p := entity.Project{}
	token := r.Header.Get("Authorization")
	var uid string
	if token != "" {
		uid, _ = authlib.ParseAuthorization(token)
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		q := project.QueryProject{}
		q.ID.Valid = true
		q.ID.Int64 = pid
		pArray, _ := project.SearchProjects(q)
		if len(pArray) != 1 {
			out.Status = "Cannot View Project Selected"
		} else {
			p = pArray[0]
			if uid != "" {
				p.Role = authlib.GetRole(uid, pid).Name
			}
		}
	}
	out.Data = p
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func putProjectsPid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	e := entity.Project{}
	p := entity.Project{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
		util.PrintDebug(out.Status)
	} else {
		ctype := r.Header.Get("Content-type")
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
			e.Description = r.FormValue("description")
			e.Visibility, err = strconv.ParseBool(r.FormValue("public"))
			if err != nil {
				util.PrintInfo(err)
				e.Visibility = false
			}
		}
		if out.Code == 0 {
			q := project.QueryProject{}
			q.ID.Valid = true
			q.ID.Int64 = pid
			pArray, _ := project.SearchProjects(q)
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
				out.Code = code.INVALIDFIELD
				out.Status = "Could not Find Matching Project"
			}
		}
		out.Data = p
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteProjectsPid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		p := entity.Project{ID: pid}
		project.DeleteProject(p)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func ProjectsPidCandidates(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCandidates(w, r)
	}
}

func getProjectsPidCandidates(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	if out.Code == 0 {
		q := candidate.CandidateQuery{}
		q.ProjectID.Int64 = pid
		q.ProjectID.Valid = true
		var err error
		out.Data, err = candidate.SearchCandidates(q)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Finding Candidates"
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func ProjectsPidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCandidatesCnid(w, r)
	case "PUT":
		putProjectsPidCandidatesCnid(w, r)
	}
}

func getProjectsPidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cnid, err := strconv.ParseInt(vars["cnid"], 10, 64)
	if out.Code == 0 && err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Candidate ID"
	}
	if out.Code == 0 {
		q := candidate.CandidateQuery{}
		q.ProjectID.Int64 = pid
		q.ProjectID.Valid = true
		q.ID.Int64 = cnid
		q.ID.Valid = true
		res, err := candidate.SearchCandidates(q)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Finding Candidates"
		} else {
			if len(res) == 1 {
				out.Data = res[0]
			} else {
				out.Code = code.UNDEFINED
				out.Status = "No Results"
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func putProjectsPidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cnid, err := strconv.ParseInt(vars["cnid"], 10, 64)
	if out.Code == 0 && err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Candidate ID"
	}

	if out.Code == 0 {
		ctype := r.Header.Get("Content-type")
		e := entity.Candidate{}
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			}
		default:
			r.ParseForm()
			traitIDStrings := strings.Split(r.FormValue("traits"), ",")
			for _, s := range traitIDStrings {
				var tid int64
				tid, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					out.Status = "Error Converting Trait ID: " + s
					out.Code = code.INVALIDFIELD
					break
				} else {
					t := entity.Trait{}
					t.ID = tid
					e.Traits = append(e.Traits, t)
				}
			}
			if out.Code == 0 {
				e.Note = r.FormValue("note")
				e.ImageID, err = strconv.ParseInt(r.FormValue("imageID"), 10, 64)
				if err != nil {
					out.Status = "Error Converting imageID to int"
					out.Code = code.INVALIDFIELD
				}
			}
		}
		if out.Code == 0 {
			e.ID = cnid
			e.ProjectID = pid
			out.Data, err = candidate.UpdateCandidate(e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Error when Updating Candidate"
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func ProjectsPidRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidRoles(w, r)
	case "POST":
		postProjectsPidRoles(w, r)
	}
}
func getProjectsPidRoles(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	if pid == 0 {
		out.Code = code.INVALIDFIELD
		out.Status = "Not allowed on Project 0 "
	}
	if out.Code == 0 {
		out.Data = authlib.GetProjectRoles(pid)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

}

func postProjectsPidRoles(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	role := authlib.Role{}
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	if pid == 0 {
		out.Code = code.INVALIDFIELD
		out.Status = "Not allowed on Project 0 "
	}
	ctype := r.Header.Get("Content-type")
	switch ctype {
	case "application/json":
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&role)
		if err != nil {
			out.Code = code.INVALIDFIELD
			out.Status = "Unable to decode json"
			util.PrintError("Unable to decode json")
		}
	default:
		r.ParseForm()
		role.UserID = r.FormValue("user_id")
		role.RoleID, err = strconv.ParseInt(r.FormValue("role_id"), 10, 64)
	}
	role.ProjectID = pid

	if out.Code == 0 {
		out.Data = role
		err = authlib.SetRole(role)
		if err != nil {
			out.Status = "Failure to assign role"
			out.Code = code.UNDEFINED
			out.Data = authlib.Role{}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

}

func ProjectsPidRolesUid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidRolesUid(w, r)
	case "PUT":
		putProjectsPidRolesUid(w, r)
	case "DELETE":
		deleteProjectsPidRolesUid(w, r)
	}
}

func getProjectsPidRolesUid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	out.Data = authlib.Role{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	uid := vars["uid"]
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Project ID must be numeric"
	}
	if pid == 0 {
		out.Code = code.INVALIDFIELD
		out.Status = "Not allowed on Project 0 "
	}
	if out.Code == 0 {
		if uid == "" {
			out.Code = code.MISSINGFIELD
			out.Status = "User ID must be set"
		}
	}
	if out.Code == 0 {
		out.Data = authlib.GetRole(uid, pid)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func putProjectsPidRolesUid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	out.Data = authlib.Role{}
	role := authlib.Role{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	ctype := r.Header.Get("Content-type")
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Project ID must be numeric"
	}
	switch ctype {
	case "application/json":
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&role)
		if err != nil {
			out.Code = code.INVALIDFIELD
			out.Status = "Unable to decode json"
			util.PrintError("Unable to decode json")
		}
	default:
		r.ParseForm()
		role.UserID = r.FormValue("user_id")
		role.RoleID, err = strconv.ParseInt(r.FormValue("role_id"), 10, 64)
	}
	role.ProjectID = pid
	role.UserID = uid
	if out.Code == 0 {
		err = authlib.UpdateRole(role)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error when updating role"
		}
		out.Data = authlib.GetRole(role.UserID, role.ProjectID)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

}
func deleteProjectsPidRolesUid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	role := authlib.Role{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Project ID must be numeric"
	}
	uid := vars["uid"]
	role.ProjectID = pid
	role.UserID = uid
	if out.Code == 0 {
		err = authlib.DeleteRole(role)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Deleting User from Project"
		}
	}
	out.Data = authlib.GetRole(role.UserID, role.ProjectID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	s := entity.Trait{Project_ID: pid}
	t, err := trait.SearchTraits(s)

	out.Data = t
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func postProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	if out.Code == 0 {
		ctype := r.Header.Get("Content-type")
		e := entity.Trait{}
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			} else {
				e.Project_ID = pid
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
			e.Type_ID, err = strconv.ParseInt(r.FormValue("type_id"), 10, 64)
			if err != nil {
				out.Code = code.INVALIDFIELD
				out.Status = "Invalid type_id"
			}
			if out.Code == 0 {
				e.Pool, err = strconv.ParseInt(r.FormValue("pool"), 10, 64)
				if err != nil {
					out.Code = code.INVALIDFIELD
					out.Status = "Invalid pool"
				}
			}
			e.Project_ID = pid
		}
		if out.Code == 0 {
			util.PrintDebug(e)
			e, err = trait.NewTrait(e)
			if err != nil {
				util.PrintDebug(err)
			}
		}
		out.Data = e
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

}

func ProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidTraitsTid(w, r)
	case "PUT":
		putProjectsPidTraitsTid(w, r)
	case "DELETE":
		deleteProjectsPidTraitsTid(w, r)
	}
}
func getProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	t := entity.Trait{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		tid, err := strconv.ParseInt(vars["tid"], 10, 64)
		if err != nil {
			out.Code = code.INVALIDFIELD
			out.Status = "Not a Numeric Trait ID"
		} else {
			s := entity.Trait{Project_ID: pid, ID: tid}
			t, err = trait.GetTrait(s)
		}
	}
	out.Data = t
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func putProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	t := entity.Trait{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		tid, err := strconv.ParseInt(vars["tid"], 10, 64)
		if err != nil {
			out.Code = code.INVALIDFIELD
			out.Status = "Not a Numeric Project ID"
		} else {
			e := entity.Trait{}
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			} else {
				e.Project_ID = pid
				e.ID = tid
				o, _ := trait.GetTrait(entity.Trait{ID: tid})
				if e.Name == "" {
					e.Name = o.Name
				}
				if e.Type_ID == 0 {
					e.Type_ID = o.Type_ID
				}
				t, _ = trait.UpdateTrait(e)
			}
		}
	}
	out.Data = t
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	t := entity.Trait{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		tid, err := strconv.ParseInt(vars["tid"], 10, 64)
		if err != nil {
			out.Code = code.INVALIDFIELD
			out.Status = "Not a Numeric Trait ID"
		} else {
			e := entity.Trait{ID: tid, Project_ID: pid}
			t, _ = trait.DeleteTrait(e)
		}
	}
	out.Data = t
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		q := cross.CrossQuery{}
		q.ProjectID.Valid = true
		q.ProjectID.Int64 = pid
		out.Data, _ = cross.SearchCrosses(q)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func postProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		ctype := r.Header.Get("Content-type")
		e := entity.Cross{}
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			} else {
				e.ProjectID = pid
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
			e.Description = r.FormValue("description")
			e.Parent1ID, err = strconv.ParseInt(r.FormValue("parent1"), 10, 64)
			if err != nil {
				out.Code = code.INVALIDFIELD
				out.Status = "Invalid parent1 field"
			}
			if out.Code == 0 {
				e.Parent2ID, err = strconv.ParseInt(r.FormValue("parent2"), 10, 64)
				if err != nil {
					out.Code = code.INVALIDFIELD
					out.Status = "Invalid parent2 field"
				}
			}
			e.ProjectID = pid
		}
		if out.Code == 0 && e.Name == "" {
			q := cross.CrossQuery{}
			q.ProjectID.Valid = true
			q.ProjectID.Int64 = pid
			curCrosses, _ := cross.SearchCrosses(q)
			count := len(curCrosses)
			e.Name = "P" + strconv.FormatInt(pid, 10) + "C" + strconv.Itoa(count)
		}
		if out.Code == 0 {
			out.Data, err = cross.CreateCross(e)
			if err != nil {
				out.Status = "Fail to add cross"
				out.Code = code.UNDEFINED
				out.Data = entity.Cross{}
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {

		cid, err := strconv.ParseInt(vars["cid"], 10, 64)
		if err != nil {
			out.Code = code.INVALIDFIELD
			out.Status = "Not a Numeric Cross ID"
		} else {
			q := cross.CrossQuery{}
			q.ProjectID.Int64 = pid
			q.ProjectID.Valid = true
			q.ID.Int64 = cid
			q.ID.Valid = true
			out.Data, err = cross.GetCross(q)
			if err != nil {
				util.PrintInfo(err)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func putProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numberic Cross ID"
	}
	if out.Code == 0 {
		ctype := r.Header.Get("Content-type")
		e := entity.Cross{}
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
			e.Description = r.FormValue("description")
			e.Parent1ID, err = strconv.ParseInt(r.FormValue("parent1"), 10, 64)
			if err != nil {
				out.Code = code.INVALIDFIELD
				out.Status = "Invalid parent1 field"
			}
			e.Parent2ID, err = strconv.ParseInt(r.FormValue("parent2"), 10, 64)
			if err != nil {
				out.Code = code.INVALIDFIELD
				out.Status = "Invalid parent2 field"
			}
			e.ProjectID = pid
		}
		if out.Code == 0 {
			if e.ID == 0 {
				e.ID = cid
			} else if e.ID != cid {
				out.Code = code.INVALIDFIELD
				out.Status = "Cross ID Mismatch"
			}
		}
		if out.Code == 0 {
			if e.ProjectID == 0 {
				e.ProjectID = pid
			} else if e.ProjectID != pid {
				out.Code = code.INVALIDFIELD
				out.Status = "Project ID Mismatch"
			}
		}
		if out.Code == 0 && e.Name == "" {
			q := cross.CrossQuery{}
			q.ID.Valid = true
			q.ID.Int64 = e.ID
			q.ProjectID.Valid = true
			q.ProjectID.Int64 = e.ProjectID
			c, _ := cross.GetCross(q)
			e.Name = c.Name
		}
		if out.Code == 0 {
			_, err = cross.UpdateCross(e)
			if err != nil {
				out.Status = "Fail to update cross"
				out.Code = code.UNDEFINED
				out.Data = entity.Cross{}
			}
		}
	}
	q := cross.CrossQuery{}
	q.ID.Valid = true
	q.ID.Int64 = cid
	q.ProjectID.Valid = true
	q.ProjectID.Int64 = pid
	out.Data, _ = cross.GetCross(q)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	e := entity.Cross{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	if out.Code == 0 {
		e.ID = cid
		e.ProjectID = pid
		cross.Delete(e)
	}
	out.Data = e
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func ProjectsPidCrossesCidPunnet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCidPunnet(w, r)
	}
}

func getProjectsPidCrossesCidPunnet(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	if out.Code == 0 {
		q := cross.CrossQuery{}
		q.ID.Valid = true
		q.ID.Int64 = cid
		q.ProjectID.Valid = true
		q.ProjectID.Int64 = pid
		me, err := cross.GetCross(q)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Failure to obtain cross info"
		}
		parents, _ := candidate.GetParents(me)
		if len(parents) > 0 {
			p1 := make(map[int64][]entity.Trait)
			p2 := make(map[int64][]entity.Trait)
			p := parents[0]
			for _, t := range p.Traits {
				p1[t.Pool] = append(p1[t.Pool], t)
			}
			if len(parents) > 1 {
				p = parents[1] // Set to other parent
			} else {
				p = parents[0] // Same parent
			}
			for _, t := range p.Traits {
				p2[t.Pool] = append(p2[t.Pool], t)
			}
			//Every pool needs 2 Traits, so if it has less, expand it
			for i, _ := range p1 {
				if len(p1[i]) == 1 {
					p1[i] = append(p1[i], p1[i][0])
				}
			}
			for i, _ := range p2 {
				if len(p2[i]) == 1 {
					p2[i] = append(p2[i], p2[i][0])
				}
			}
			type Chance struct {
				Trait        entity.Trait `json:"trait"`
				PercentCarry float64      `json:"carry"`
				PercentShow  float64      `json:"show"`
				Total        int          `json:"count"`
			}
			type Pair struct {
				M entity.Trait
				F entity.Trait
			}
			output := make(map[int64]*Chance)
			for i, p := range p1 { // for each pool
				var pairs []Pair
				for _, f := range p { // for each trait in the pool
					for _, m := range p2[i] { // for each trait in the p2 pool
						output[m.ID] = &Chance{Trait: m, PercentCarry: 0, PercentShow: 0}
						output[f.ID] = &Chance{Trait: f, PercentCarry: 0, PercentShow: 0}
						pairs = append(pairs, Pair{M: m, F: f})
					}
				}
				for _, pair := range pairs {
					m := pair.M
					f := pair.F
					util.PrintDebug(m.Name + " vs " + f.Name)
					o := output[m.ID]
					o.PercentCarry += 1
					if f.ID != m.ID {
						o = output[f.ID]
						o.PercentCarry += 1
						if f.Type_ID == 2 {
							o = output[f.ID]
							o.PercentShow += 1
						}
						if m.Type_ID == 2 {
							o = output[m.ID]
							o.PercentShow += 1
						}
						if m.Type_ID == 1 && f.Type_ID == 1 {
							o = output[m.ID]
							o.PercentShow += 1
							o = output[f.ID]
							o.PercentShow += 1
						}
					} else {
						o = output[m.ID]
						o.PercentShow += 1
					}
				}
				for _, e := range output {
					e.Total = len(pairs)
				}
			}
			cList := make([]*Chance, 0)
			for _, e := range output {
				cList = append(cList, e)
			}
			out.Data = cList
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)

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
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	if out.Code == 0 {
		q := candidate.CandidateQuery{}
		q.ProjectID.Valid = true
		q.ProjectID.Int64 = pid
		q.CrossID.Valid = true
		q.CrossID.Int64 = cid
		out.Data, _ = candidate.SearchCandidates(q)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func postProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}

	if out.Code == 0 {
		ctype := r.Header.Get("Content-type")
		e := entity.Candidate{}
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			}
		default:
			r.ParseForm()
			traitIDStrings := strings.Split(r.FormValue("traits"), ",")
			for _, s := range traitIDStrings {
				var tid int64
				tid, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					out.Status = "Error Converting Trait ID: " + s
					out.Code = code.INVALIDFIELD
					break
				} else {
					t := entity.Trait{}
					t.ID = tid
					e.Traits = append(e.Traits, t)
				}
			}
			if out.Code == 0 {
				e.Note = r.FormValue("note")
				e.ImageID, err = strconv.ParseInt(r.FormValue("imageID"), 10, 64)
				if err != nil {
					out.Status = "Error Converting imageID to int"
					out.Code = code.INVALIDFIELD
				}
			}
		}

		if out.Code == 0 {
			e.ProjectID = pid
			e.CrossID = cid
			out.Data, _ = candidate.CreateCandidate(e)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
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
	//token := r.Header.Get("Authorization")
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	cnid, err := strconv.ParseInt(vars["cnid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Candidate ID"
	}
	if out.Code == 0 {
		q := candidate.CandidateQuery{}
		q.ProjectID.Valid = true
		q.ProjectID.Int64 = pid
		q.CrossID.Valid = true
		q.CrossID.Int64 = cid
		q.ID.Valid = true
		q.ID.Int64 = cnid
		cn, err := candidate.SearchCandidates(q)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Searching for Candidate"
		}
		if len(cn) == 1 {
			out.Data = cn[0]
		} else {
			out.Data = entity.Candidate{}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func putProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	cnid, err := strconv.ParseInt(vars["cnid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Candidate ID"
	}

	if out.Code == 0 {
		ctype := r.Header.Get("Content-type")
		e := entity.Candidate{}
		switch ctype {
		case "application/json":
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Invalid JSON Posted"
				util.PrintError("Unable to decode json")
				util.PrintDebug(err)
			}
		default:
			r.ParseForm()
			traitIDStrings := strings.Split(r.FormValue("traits"), ",")
			for _, s := range traitIDStrings {
				var tid int64
				tid, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					out.Status = "Error Converting Trait ID: " + s
					out.Code = code.INVALIDFIELD
					break
				} else {
					t := entity.Trait{}
					t.ID = tid
					e.Traits = append(e.Traits, t)
				}
			}
			if out.Code == 0 {
				e.Note = r.FormValue("note")
				e.ImageID, err = strconv.ParseInt(r.FormValue("imageID"), 10, 64)
				if err != nil {
					out.Status = "Error Converting imageID to int"
					out.Code = code.INVALIDFIELD
				}
			}
		}
		if out.Code == 0 {
			e.ID = cnid
			e.ProjectID = pid
			e.CrossID = cid
			out.Data, err = candidate.UpdateCandidate(e)
			if err != nil {
				out.Code = code.UNDEFINED
				out.Status = "Error when Updating Candidate"
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func deleteProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	cnid, err := strconv.ParseInt(vars["cnid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Candidate ID"
	}

	if out.Code == 0 {
		q := entity.Candidate{}
		q.ProjectID = pid
		q.CrossID = cid
		q.ID = cnid
		out.Data, err = candidate.DeleteCandidate(q)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Deleting Candidate"
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func ProjectsPidTreview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsTreview(w, r)
	}
}
func getProjectsTreview(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	if out.Code == 0 {
		var err error
		out.Data, err = treview.GenerateForest(pid)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Generating Treenode objects"
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func ProjectsPidTreviewCid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsTreviewCid(w, r)
	}
}
func getProjectsTreviewCid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	}
	cid, err := strconv.ParseInt(vars["cid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Cross ID"
	}
	if out.Code == 0 {
		var err error
		out.Data, err = treview.Generate(pid, cid)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error Generating Treenode objects"
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
