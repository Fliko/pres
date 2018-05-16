package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type taterbase []string

var messages taterbase

func sayHello(w http.ResponseWriter, r *http.Request) {
	var msg string

	for _, m := range messages {
		msg += m + "\n"
	}
	w.Write([]byte(msg))
}

func store(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]
	messages = append(messages, message)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sayHello)
	router.HandleFunc("/{msg}", store)

	log.Fatal(http.ListenAndServe(":8080", router))
}
