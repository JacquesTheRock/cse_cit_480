package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"treview.com/bloom/entity"
)

func ProjectsRoot(w http.ResponseWriter, r *http.Request) {
	c := entity.Cross{ID: 0, ProjectID: 0, Name: "Bobby"}
	json.NewEncoder(w).Encode(c)
	fmt.Fprintln(w, "This is the Projects Root API")
}
