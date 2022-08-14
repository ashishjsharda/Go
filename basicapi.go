package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome !!")
	fmt.Println("Endpoint Hit: greeting")
}
func handleRequests() {
	http.HandleFunc("/", greeting)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
