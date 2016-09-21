package handlers

import (
	"net/http"
)
func Users(w http.ResponseWriter, r *http.Request) { }
func UsersUid(w http.ResponseWriter, r *http.Request) { }
func UsersUidProjects(w http.ResponseWriter, r *http.Request) { }
func UsersUidMail(w http.ResponseWriter, r *http.Request) { }
func UsersUidMailMid(w http.ResponseWriter, r *http.Request) { }
func Projects(w http.ResponseWriter, r *http.Request) { }
func ProjectsPid(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidTraits(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidCrosses(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidTreview(w http.ResponseWriter, r *http.Request) { }
func ProjectsPidTreviewCid(w http.ResponseWriter, r *http.Request) { }
func Auth(w http.ResponseWriter, r *http.Request) { }
func Breeds(w http.ResponseWriter, r *http.Request) { }
func BreedsBid(w http.ResponseWriter, r *http.Request) { }
func BreedsBidTraits(w http.ResponseWriter, r *http.Request) { }
func BreedsBidTraitsTid(w http.ResponseWriter, r *http.Request) { }
