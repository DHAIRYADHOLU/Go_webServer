package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	}

	// Creating server and specific port to run
	server := http.Server{
		Addr:    ":8080", // port no
		Handler: http.HandlerFunc(handler),
	}

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
