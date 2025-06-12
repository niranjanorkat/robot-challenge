package main

import "fmt"

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ArgCheck defines a function that validates a single string argument and returns a typed result or error.
type ArgCheck func(string) (interface{}, error)

// checkParts applies a list of validation functions to corresponding input parts.
// It returns the parsed results or false with an error message if any check fails.
func checkArgs(parts []string, checks []ArgCheck, errMsg string) ([]interface{}, bool) {
	if len(parts) < len(checks) {
		fmt.Println(errMsg)
		return nil, false
	}
	results := make([]interface{}, len(checks))
	for i, check := range checks {
		val, err := check(parts[i])
		if err != nil {
			fmt.Println(errMsg)
			return nil, false
		}
		results[i] = val
	}
	return results, true
}

// checkWarehouseID validates and extracts the warehouse ID from a string like "W1"
func checkWarehouseID(warehousesLen int) ArgCheck {
	return func(s string) (interface{}, error) {
		id, err := strconv.Atoi(strings.TrimPrefix(s, "W"))
		if err != nil || id <= 0 || id > warehousesLen {
			return nil, fmt.Errorf("invalid warehouse ID")
		}
		return id, nil
	}
}

// checkCoordinate validates and parses a coordinate integer within bounds.
func checkCoordinate(max int) ArgCheck {
	return func(s string) (interface{}, error) {
		n, err := strconv.Atoi(s)
		if err != nil || n < 0 || n >= max {
			return nil, fmt.Errorf("invalid coordinate")
		}
		return n, nil
	}
}

func checkArgsLength(parts []string, expected int, errMsg string) bool {
	if len(parts) != expected {
		fmt.Println(errMsg)
		return false
	}
	return true
}