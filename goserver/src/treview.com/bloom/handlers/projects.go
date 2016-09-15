package handlers

import (
	"fmt"
	"net/http"
)

func ProjectsRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"This is the Projects Root API")
}
