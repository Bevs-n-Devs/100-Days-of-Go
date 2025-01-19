package main

import "fmt"

func Example4() {
	fmt.Println("\nExample 4: Simplicity in Concurrency")

	var scenario = `
	Scenario:
	You need to fetch data from two APIs concurrently.

	Over-Complicated Approach:
	Using overly complex sychronization mechanisms.

	Code Example:
	func fetchDataFromAPIs() ([]string, error) {
		api1Data := make(chan string)
		api2Data := make(chan string)
		errors := make(chan error, 2)

		go func() {
			data, err := fetchFromAPI1()
			if err != nil {
				errors <- err
				return
			}
			api1Data <- data
		}()

		go func() {
			data, err := fetchFromAPI2()
			if err != nil {
				errors <- err
				return
			}
			api2Data <- data
		}()

		var results []string
		for i := 0; i < 2; i++ {
			select {
			case data := <-api1Data:
				results = append(results, data)
			case data := <-api2Data:
				results = append(results, data)
			case err := <-errors:
				return nil, err
			}
		}

		return results, nil
	}

	Problems:
		- Too many channels and select statements.
		- Hard to follow logic.

	KISS Approach:
	Simplify usign sync.WaitGroup.

	Code Example:
	func fetchDataFromAPIs() ([]string, error) {
		var wg sync.WaitGroup
		var mu sync.Mutex
		var results []string
		var err error

		wg.Add(2)

		go func() {
			defer wg.Done()
			data, e := fetchFromAPI1()
			mu.Lock()
			defer mu.Unlock()
			if e != nil {
				err = e
				return
			}
			results = append(results, data)
		}()

		go func() {
			defer wg.Done()
			data, e := fetchFromAPI2()
			mu.Lock()
			defer mu.Unlock()
			if e != nil {
				err = e
				return
			}
			results = append(results, data)
		}()

		wg.Wait()
		return results, err
	}

	Benefits:
		- Fewer channels.
		- Easier to understand and maintain.
	`
	fmt.Println(scenario)
}
