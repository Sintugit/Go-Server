package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post requested sucessful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	filesServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", filesServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8181")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal(err)
	}
}
