package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", router))
}