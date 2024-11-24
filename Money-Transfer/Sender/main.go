package main

import "fmt"

func main() {

	// URL where the receiver service is running
	receiverURL := "http://localhost:8080/receive"
	apiToken := "secure-token-12345" // Example tokentoken
	// Prepare transfer data

	transfer := TransferData{
		FromAccount: "AccountA",
		ToAccount:   "AccountB",
		Amount:      1000.00,
	}

	// Initiate the transfer
	if err := initiateTransfer(receiverURL, apiToken, transfer); err != nil {
		fmt.Println("Error during transfer:", err)
	}
}
