/*
Problem Statement #10: Configuration Validator

Objective:
	Write a Go program with a function that validates a configuration map (map[string]string) and returns a structured, custom error if the validation fails.

Rules/Logic:

	1. Define a custom error struct named ConfigError with two fields: Field (string) and Reason (string).

	2. Implement the error interface for your ConfigError struct. The Error() method should produce a message like "Configuration Error: Field 'PORT' is missing".

	3. Create a function validateConfig(config map[string]string) error. This function returns a *ConfigError if validation fails, and nil if it succeeds.

	4. The validateConfig function must check the following rules:

		1. The key "PORT" must exist. If not, return a ConfigError.

		2. The key "ENV" must exist. If not, return a ConfigError.

		3. The value for the "ENV" key must be one of "development", "staging", or "production". If it's anything else, return a ConfigError.

	5. In your main function, create several test cases (a valid map, a map with a missing key, a map with an invalid value) and print the result of calling validateConfig for each.

---Terminal Output---
	// Testcase 1: Valid config
	// Testcase 2: Missing PORT
	// Testcase 3: Missing ENV
	// Testcase 4: Invalid ENV value
*/

package main

import "fmt"

type ConfigError struct {
	Field  string
	Reason string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("Configuration error : Field %s %s", e.Field, e.Reason)
}

func validateConfig(config map[string]string) error {
	if _, ok := config["PORT"]; !ok {
		return &ConfigError{Field: "PORT", Reason: "key is missing in config map"}
	}
	if _, ok := config["ENV"]; !ok {
		return &ConfigError{Field: "ENV", Reason: "key is missing in config map"}
	}
	if val, ok := config["ENV"]; ok {
		if val != "development" && val != "staging" && val != "production" {
			return &ConfigError{Field: "ENV", Reason: "values are not matching in ENV key in config map"}
		}
	}
	return nil
}

func main() {
	fmt.Println("code logic here")
	// Test Case 1: A valid config map
	fmt.Println("--- Testing Valid Config ---")
	validConfig := map[string]string{
		"PORT": "8080",
		"ENV":  "production",
	}
	err := validateConfig(validConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config is valid.")
	}

	// Test Case 2: A map with a missing key
	fmt.Println("\n--- Testing Missing Key ---")
	missingKeyConfig := map[string]string{
		"ENV": "development",
	}
	err = validateConfig(missingKeyConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config is valid.")
	}

	// Test Case 3: A map with an invalid value
	fmt.Println("\n--- Testing Invalid Value ---")
	invalidValueConfig := map[string]string{
		"PORT": "8080",
		"ENV":  "testing",
	}
	err = validateConfig(invalidValueConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config is valid.")
	}
}
