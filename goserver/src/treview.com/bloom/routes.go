package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"treview.com/bloom/handlers"
)

type Route struct {
	Name string
	Methods []string
	Pattern string
	HandlerFunc http.HandlerFunc
}


func NewRouter(root string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := route.HandlerFunc
		router.
			Methods(route.Methods...).
			Path(root + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

type Routes []Route
var routes = Routes {
	Route {
		"users",
		[]string{"GET","POST"},
		"/users",
		handlers.Users,
	},
	Route {
		"users_uid",
		[]string{"GET","PUT"},
		"/users/{uid}",
		handlers.UsersUid,
	},
	Route {
		"users_uid_projects",
		[]string{"GET"},
		"/users/{uid}/projects",
		handlers.UsersUidProjects,
	},
	Route {
		"users_uid_mail",
		[]string{"GET","POST"},
		"/users/{uid}/mail",
		handlers.UsersUidMail,
	},
	Route {
		"users_uid_mail_mid",
		[]string{"GET","PUT","DELETE"},
		"/users/{uid}/mail/{mid}",
		handlers.UsersUidMailMid,
	},
	Route {
		"projects",
		[]string{"GET","POST"},
		"/projects",
		handlers.Projects,
	},
	Route {
		"projects_pid",
		[]string{"GET","PUT","DELETE"},
		"/projects/{pid}",
		handlers.ProjectsPid,
	},
	Route {
		"projects_pid_traits",
		[]string{"GET","POST"},
		"/projects/{pid}/traits",
		handlers.ProjectsPidTraits,
	},
	Route {
		"projects_pid_traits_tid",
		[]string{"GET","PUT","DELETE"},
		"/projects/{pid}/traits/{tid}",
		handlers.ProjectsPidTraitsTid,
	},
	Route {
		"projects_pid_crosses",
		[]string{"GET","POST"},
		"/projects/{pid}/crosses",
		handlers.ProjectsPidCrosses,
	},
	Route {
		"projects_pid_crosses_cid",
		[]string{"GET","PUT","DELETE"},
		"/projects/{pid}/crosses/{cid}",
		handlers.ProjectsPidCrossesCid,
	},
	Route {
		"projects_pid_crosses_cid_candidates",
		[]string{"GET","POST"},
		"/projects/{pid}/crosses/{cid}/candidates",
		handlers.ProjectsPidCrossesCidCandidates,
	},
	Route {
		"projects_pid_crosses_cid_candidates_cnid",
		[]string{"GET","PUT","DELETE"},
		"/projects/{pid}/crosses/{cid}/candidates/{cnid}",
		handlers.ProjectsPidCrossesCidCandidatesCnid,
	},
	Route {
		"projects_pid_treview",
		[]string{"GET"},
		"/projects/{pid}/treview",
		handlers.ProjectsPidTreview,
	},
	Route {
		"projects_pid_treview_cid",
		[]string{"GET"},
		"/projects/{pid}/treview/{cid}",
		handlers.ProjectsPidTreviewCid,
	},
	Route {
		"auth",
		[]string{"GET","POST","DELETE"},
		"/auth",
		handlers.Auth,
	},
	Route {
		"breeds",
		[]string{"GET","POST"},
		"/breeds",
		handlers.Breeds,
	},
	Route {
		"breeds_bid",
		[]string{"GET","PUT"},
		"/breeds/{bid}",
		handlers.BreedsBid,
	},
	Route {
		"breeds_bid_traits",
		[]string{"GET","POST"},
		"/breeds/{bid}/traits",
		handlers.BreedsBidTraits,
	},
	Route {
		"breeds_bid_traits_tid",
		[]string{"GET","PUT","DELETE"},
		"/breeds/{bid}/traits/{tid}",
		handlers.BreedsBidTraitsTid,
	},
}
