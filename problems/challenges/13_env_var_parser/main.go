package main

import (
	"errors"
	"fmt"
	"strings"
)

func parseEnvVars(envVars []string) (map[string]string, error) {
	var result = make(map[string]string)
	for _, word := range envVars {
		parts := strings.SplitN(word, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("malformed entry: " + word)
		}
		result[parts[0]] = parts[1]
	}
	return result, nil
}

func main() {
	vars := []string{
		"PORT=8080",
		"DB_USER=root",
		"DB_PASS=password123",
		"INVALID_VAR", // A malformed entry
	}

	config, err := parseEnvVars(vars)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed Config: %v\n", config)
	}

	// Also test a valid case
	varsValid := []string{
		"HOST=localhost",
		"TIMEOUT=30s",
	}

	configValid, err := parseEnvVars(varsValid)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed Config: %v\n", configValid)
	}
}
