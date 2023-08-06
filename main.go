package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Lets go"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("Starting server on port :8080!")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
