package main

import (
	"Learn-golang-backend/handlers"
	"Learn-golang-backend/middleware"
	"log"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// server port
const (
	WEBSERVERPORT = ":8080"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler).Methods("GET")

	// http.Handle("/", r)
	// http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))
	log.Fatal(http.ListenAndServe(WEBSERVERPORT, nil))
}
