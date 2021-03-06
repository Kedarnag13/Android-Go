package main

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Android-Go/api/v1/controllers/account"
	"github.com/kedarnag13/Android-Go/api/v1/controllers/users"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", account.Registration.Create).Methods("POST")
	r.HandleFunc("/login", account.Session.Create).Methods("POST")
	r.HandleFunc("/logout/{devise_token:([a-zA-Z0-9]+)?}", account.Session.Destroy).Methods("GET")

	// Send SMS Invite
	r.HandleFunc("/invite", users.Message.Send).Methods("POST")

	http.Handle("/", r)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3010 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3010", nil))
}
