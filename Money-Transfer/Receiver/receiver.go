package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// TransferData represents the structure of the incoming transfer data.
type TransferData struct {
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Amount      float64 `json:"amount"`
}

// Simulated account balances (protected by mutex)

var accountBalances = map[string]float64{
	"AccountA": 5000.0,
	"AccountB": 3000.0,
}
var mutex = &sync.Mutex{}

const validToken = "secure-token-12345"

// HandleReceive process the transfer request and updates the receiver's records.
func HandleReceive(w http.ResponseWriter, r *http.Request) {

	//Check HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Authenticate the request
	authHeader := r.Header.Get("Authorization")
	if authHeader != "Bearer "+validToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode the incoming transfer data
	var transfer TransferData
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Validate transfer data
	if transfer.Amount <= 0 || transfer.FromAccount == "" || transfer.ToAccount == "" {
		http.Error(w, "Invalid transfer data", http.StatusBadRequest)
		return
	}

	// Process the transfer atomically
	mutex.Lock()
	defer mutex.Unlock()

	fromBalance, fromExists := accountBalances[transfer.FromAccount]
	_, toExists := accountBalances[transfer.ToAccount]

	if !fromExists || !toExists {
		http.Error(w, "Account does not exist", http.StatusBadRequest)
		return
	}

	if fromBalance < transfer.Amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	// Perfrom the transfer
	accountBalances[transfer.FromAccount] -= transfer.Amount
	accountBalances[transfer.ToAccount] += transfer.Amount

	// Log the successfull transaction
	log.Printf("Transfer completed: %+v\n", transfer)

	// Respond with success

	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Transfer successful",
	}
	_ = json.NewEncoder(w).Encode(response)

}
