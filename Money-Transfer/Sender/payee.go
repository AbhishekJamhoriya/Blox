package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// trasferdata holds the infromation for a money trasfer.
type TransferData struct {
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Amount      float64 `json:"amount"`
}

// InitialeTransfer sends the transfer request to the receiver.

func initiateTransfer(receiverURL, apiToken string, trasfer TransferData) error {
	// Validate input data
	if trasfer.Amount <= 0 {
		return fmt.Errorf("invalid transfer amount: %f", trasfer.Amount)
	}
	if trasfer.FromAccount == "" || trasfer.ToAccount == "" {
		return fmt.Errorf("amount details are incomplete")
	}

	// Marshal the transfer data into JSON
	data, err := json.Marshal(trasfer)
	if err != nil {
		return fmt.Errorf("error encoding transfer data: %w", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Create the request with authentication token
	req, err := http.NewRequest("POST", receiverURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error initiatingtransfer data: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	// Retry mechanism
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Attempt %d failed: %s\n", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}
		defer resp.Body.Close()
		//Check if the transfer was successfull
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Transfer successfull")
			return nil

		} else {
			return fmt.Errorf("transfer failed, status: %s", resp.Status)
		}

	}

	return fmt.Errorf("transfer failed after %d retries", maxRetries)
}
