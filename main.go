package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/user", userHandler)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello world"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the about page</h1>"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	if userId == "" {
		http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := map[string]string{"userId": userId}
	jsonResp, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
