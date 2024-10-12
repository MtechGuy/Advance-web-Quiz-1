package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Set handler for the root URL "/"

	fmt.Println("Starting server at :3000")
	err := http.ListenAndServe(":3000", nil) // Listen on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
