package main

import "fmt"

func Example3() {
	fmt.Println("\nExample 3: Over-Engineering Configuration")

	var scenario = `
	Scenario:
	You need to configure an application with default values.

	Over-Complicated Approach:
	Using a highly flexible but overly complex system.

	Code Example:
	type Config struct {
		Port        int
		Environment string
	}

	func LoadConfig() (*Config, error) {
		port, err := strconv.Atoi(os.Getenv("APP_PORT"))
		if err != nil {
			port = 8080 // default port
		}

		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "production" // default environment
		}

		return &Config{Port: port, Environment: env}, nil
	}

	Problems:
		1. Complex environment variable parsing.
		2. Extra error handling when defaults would suffice.

	KISS Approach:
	Use hardcoded defaults with simple overrides.

	Code Example:
	type Config struct {
		Port        int
		Environment string
	}

	func DefaultConfig() *Config {
		return &Config{
			Port:        8080,
			Environment: "production",
		}
	}
	
	Benefits:
		- Simpler logic.
		- Easier to understand, maintain, and test.
	`
	fmt.Println(scenario)
}
