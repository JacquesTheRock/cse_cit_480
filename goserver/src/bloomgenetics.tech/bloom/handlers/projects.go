package handlers

import (
	authlib "bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/code"
	"bloomgenetics.tech/bloom/cross"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/project"
	"bloomgenetics.tech/bloom/trait"
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
	out := entity.ApiData{}
	p, _ := project.SearchProjects(entity.Project{})
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
			e.Visibility, err = strconv.ParseBool(r.FormValue("public"))
			if err != nil {
				util.PrintError(err)
				e.Visibility = false
			}
		}
		if e.Name != "" {
			out.Data, err = project.NewProject(uid, e)
		} else {
			out.Code = code.MISSINGFIELD
			out.Status = "Projects must have a name"
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
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a Numeric Project ID"
	} else {
		pArray, _ := project.SearchProjects(entity.Project{ID: pid})
		if len(pArray) != 1 {
			out.Status = "Cannot View Project Selected"
		} else {
			p = pArray[0]
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
				util.PrintError(err)
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
		if out.Code != 0 {
			pArray, _ := project.SearchProjects(entity.Project{ID: pid})
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
	if out.Code != 0 {
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
				util.PrintError(err)
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
			} else {
				e.Project_ID = pid
			}
		}
		if e.Project_ID == pid {
			e, _ = trait.NewTrait(e)
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
				util.PrintError(err)
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
				util.PrintError(err)
			} else {
				e.ProjectID = pid
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
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
		if out.Code == 0 && e.Name == "" {
			q := cross.CrossQuery{}
			q.ProjectID.Valid = true
			q.ProjectID.Int64 = pid
			curCrosses, _ := cross.SearchCrosses(q)
			count := len(curCrosses)
			e.Name = "PROJECT" + strconv.FormatInt(pid, 10) + "CROSS" + strconv.Itoa(count)
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
				util.PrintError(err)
			}
		default:
			r.ParseForm()
			e.Name = r.FormValue("name")
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

func ProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCidCandidates(w, r)
	case "POST":
		postProjectsPidCrossesCidCandidates(w, r)
	}
}
func getProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	//token := r.Header.Get("Authorization")
	c := [10]entity.Candidate{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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
	//token := r.Header.Get("Authorization")
	c := entity.Candidate{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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
