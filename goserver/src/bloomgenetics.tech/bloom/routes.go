package main

import (
	"bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/handlers"
	"bloomgenetics.tech/bloom/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Wrapper(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			util.PrintInfo(
				r.Method + "\t" +
					r.RequestURI + "\t" +
					name + "\t")

			vars := mux.Vars(r)
			var err error
			var pid int64
			pid = 0
			if vars["pid"] != "" {
				pid, err = strconv.ParseInt(vars["pid"], 10, 64)
			}
			if err != nil {
				util.PrintDebug(err)
				util.PrintInfo("Failed to parse pid in Wrapper")
				pid = 0
			}
			token := r.Header.Get("Authorization")
			uid, _ := auth.ParseAuthorization(token)
			if uid == "" {
				uid = "guest"
			}
			res, err := auth.CheckAuth(uid, pid, name, r.Method)
			util.PrintDebug("Auth against: " + r.Method + " " + name + " result: " + strconv.FormatBool(res))
			if res {
				inner.ServeHTTP(w, r)
			} else {
				var unAuth http.HandlerFunc
				unAuth = handlers.UnAuthorized
				unAuth.ServeHTTP(w, r)
			}
		})
}

func NewRouter(root string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := Wrapper(route.HandlerFunc, route.Name)
		router.
			Methods(route.Methods...).
			Path(root + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

type Routes []Route

var routes = Routes{
	Route{
		"users",
		[]string{"GET", "POST"},
		"/users",
		handlers.Users,
	},
	Route{
		"users_uid",
		[]string{"GET", "PUT"},
		"/users/{uid}",
		handlers.UsersUid,
	},
	Route{
		"users_uid_projects",
		[]string{"GET"},
		"/users/{uid}/projects",
		handlers.UsersUidProjects,
	},
	Route{
		"users_uid_projects_pid",
		[]string{"GET", "DELETE"},
		"/users/{uid}/projects/{pid}",
		handlers.UsersUidProjectsPid,
	},
	Route{
		"users_uid_mail",
		[]string{"GET", "POST"},
		"/users/{uid}/mail",
		handlers.UsersUidMail,
	},
	Route{
		"users_uid_mail_mid",
		[]string{"GET", "PUT", "DELETE"},
		"/users/{uid}/mail/{mid}",
		handlers.UsersUidMailMid,
	},
	Route{
		"projects",
		[]string{"GET", "POST"},
		"/projects",
		handlers.Projects,
	},
	Route{
		"projects_pid",
		[]string{"GET", "PUT", "DELETE"},
		"/projects/{pid}",
		handlers.ProjectsPid,
	},
	Route{
		"projects_pid_traits",
		[]string{"GET", "POST"},
		"/projects/{pid}/traits",
		handlers.ProjectsPidTraits,
	},
	Route{
		"projects_pid_traits_tid",
		[]string{"GET", "PUT", "DELETE"},
		"/projects/{pid}/traits/{tid}",
		handlers.ProjectsPidTraitsTid,
	},
	Route{
		"projects_pid_crosses",
		[]string{"GET", "POST"},
		"/projects/{pid}/crosses",
		handlers.ProjectsPidCrosses,
	},
	Route{
		"projects_pid_crosses_cid",
		[]string{"GET", "PUT", "DELETE"},
		"/projects/{pid}/crosses/{cid}",
		handlers.ProjectsPidCrossesCid,
	},
	Route{
		"projects_pid_crosses_cid_candidates",
		[]string{"GET", "POST"},
		"/projects/{pid}/crosses/{cid}/candidates",
		handlers.ProjectsPidCrossesCidCandidates,
	},
	Route{
		"projects_pid_crosses_cid_candidates_cnid",
		[]string{"GET", "PUT", "DELETE"},
		"/projects/{pid}/crosses/{cid}/candidates/{cnid}",
		handlers.ProjectsPidCrossesCidCandidatesCnid,
	},
	Route{
		"projects_pid_treview",
		[]string{"GET"},
		"/projects/{pid}/treview",
		handlers.ProjectsPidTreview,
	},
	Route{
		"projects_pid_treview_cid",
		[]string{"GET"},
		"/projects/{pid}/treview/{cid}",
		handlers.ProjectsPidTreviewCid,
	},
	Route{
		"auth",
		[]string{"GET", "POST", "DELETE"},
		"/auth",
		handlers.Auth,
	},
	Route{
		"projects_pid_roles",
		[]string{"GET", "POST"},
		"/projects/{pid}/roles",
		handlers.ProjectsPidRoles,
	},
	Route{
		"projects_pid_roles_uid",
		[]string{"GET", "PUT", "DELETE"},
		"/projects/{pid}/roles/{uid}",
		handlers.ProjectsPidRolesUid,
	},
}
