package main

import (
	"encoding/json"
	"math/big"
	"strings"
)

// Parse a JSON string into Go data structures with arbitrary precision
func parseJSON(jsonStr string) (interface{}, error) {
	var result interface{}

	//Use a JSON decoder with custom handling for numbers
	decoder := json.NewDecoder(strings.NewReader(jsonStr))
	decoder.UseNumber() //Ensure numbers are parsed as json.Number

	//Decode the JSON string

	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	//Convert numbers to arbitrary precision values
	return convertToArbitraryPrecision(result), nil
}

// Recursively convert json.Number to arbitrary precision integers or floats

func convertToArbitraryPrecision(value interface{}) interface{} {
	switch v := value.(type) {
	case json.Number: //if value is a number
		if i, err := v.Int64(); err == nil {
			return big.NewInt(i) //convert to big.Int
		}
		if f, err := v.Float64(); err == nil {
			return big.NewFloat(f) //convert to big.Float
		}
	case map[string]interface{}: //if value is a map
		for key, val := range v {
			v[key] = convertToArbitraryPrecision(val) //Recursively convert values
		}
	case []interface{}: //if value is a list
		for i, val := range v {
			v[i] = convertToArbitraryPrecision(val) // recursively convert elements
		}
	}
	return value // Return value if no conversion is needed
}
