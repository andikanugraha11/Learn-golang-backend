package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Function for handle profile
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.Write([]byte(username))
}
