package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Handle the transfer request
	http.HandleFunc("/receive", HandleReceive)

	// Start the server
	fmt.Println("Receive running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
