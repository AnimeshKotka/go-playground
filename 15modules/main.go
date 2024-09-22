package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go modules")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":2000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
