package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello, world!")
	// d, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, "Bad request! Error: %v", err)
	// }
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", helloHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
