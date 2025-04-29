package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: "1", Name: "John Doe", Email: "john.doe@example.com"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	port := ":8888"
	http.HandleFunc("/user", getUserHandler)
	log.Println("REST server listening on " + port)
	http.ListenAndServe(port, nil)
}
