package main

import (
	"fmt"
	"log"
	"net/http"
)

type form struct {
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Invalid Method Used", http.StatusNotFound)
	// }

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return

	}
	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello there")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
