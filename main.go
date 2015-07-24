package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", account.Registration.Create)
	http.Handle("/", r)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3010 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3010", nil))
}
