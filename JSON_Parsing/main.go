package main

import (
	"fmt"
)

func main() {
	jsonStr := `{"number": "12345678901234567890.123456789", "list": [1, 2, 3], "map": {"key": "value"}}`

	parsed, err := parseJSON(jsonStr)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}
	fmt.Printf("Parsed Data: %#v\n", parsed)
}
