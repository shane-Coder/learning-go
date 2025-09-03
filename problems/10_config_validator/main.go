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

// Improvement: Define constants for magic strings
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

type ConfigError struct {
	Field  string
	Reason string
}

func (e *ConfigError) Error() string {
	// Improvement: Use single quotes for clarity
	return fmt.Sprintf("Configuration Error: Field '%s' %s", e.Field, e.Reason)
}

func validateConfig(config map[string]string) error {
	if _, ok := config["PORT"]; !ok {
		return &ConfigError{Field: "PORT", Reason: "is missing"}
	}

	// Improvement: More efficient check
	env, ok := config["ENV"]
	if !ok {
		return &ConfigError{Field: "ENV", Reason: "is missing"}
	}

	// Check the value against our constants
	switch env {
	case EnvDevelopment, EnvStaging, EnvProduction:
		// Valid case, do nothing.
	default:
		// Invalid case.
		return &ConfigError{Field: "ENV", Reason: fmt.Sprintf("has an invalid value: '%s'", env)}
	}

	return nil
}

func main() {
	fmt.Println("--- Validating Configurations ---")

	// Improvement: Use a table-driven test structure
	testCases := []struct {
		name   string // A name for the test case
		config map[string]string
	}{
		{
			name: "Valid Config",
			config: map[string]string{
				"PORT": "8080",
				"ENV":  EnvProduction,
			},
		},
		{
			name: "Missing PORT Key",
			config: map[string]string{
				"ENV": EnvDevelopment,
			},
		},
		{
			name: "Invalid ENV Value",
			config: map[string]string{
				"PORT": "3000",
				"ENV":  "testing",
			},
		},
		{
			name: "Missing ENV Key",
			config: map[string]string{
				"PORT": "9000",
			},
		},
	}

	// Loop over the test cases
	for _, tc := range testCases {
		fmt.Printf("\nRunning Test: '%s'\n", tc.name)
		err := validateConfig(tc.config)
		if err != nil {
			fmt.Printf("Result: FAILED - %v\n", err)
		} else {
			fmt.Println("Result: PASSED - Config is valid.")
		}
	}
}
