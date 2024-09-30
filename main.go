package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Building a formula 1 App!")

}
