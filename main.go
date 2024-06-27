package main

import (
	"fmt"
	"log"
	"net/http"
)


func formHandler (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return;
	}

	fmt.Fprintf(w, "Post Request Successful!")

	name := r.FormValue("name")

	fmt.Fprintf(w, "name is: %s", name)
}

func helloHandler (w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found!", http.StatusNotFound)
		return
	} 

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./src/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("server is listening to 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}