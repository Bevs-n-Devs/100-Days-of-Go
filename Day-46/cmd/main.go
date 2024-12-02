package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nUnit Testing in Go")

	objectives := `
1. Understand the basics of unit testing in Go.
2. Learn how to structure and write tests.
3. Explore table-driven testing (a Go best practice).
4. Learn to use mocks and the testify library for advanced testing.
5. Understand best practices in testing.
	`
	fmt.Println(objectives)

	bestPractices := `
	Best Practices

1. Write Clear Test Cases:
	- Test one specific behavior per test function.
2.Use Table-Driven Tests:
	- Simplifies testing multiple scenarios.
3. Name Tests Meaningfully:
	- Include the function name and purpose (e.g., TestAdd_TwoPositives).
4. Mock External Dependencies:
	- Avoid testing dependencies like databases or external APIs in unit tests.
5. Keep Tests Independent:
	- Avoid shared state across tests.
6. Test Edge Cases:
	- Cover all scenarios, including invalid inputs.
	`
	fmt.Println(bestPractices)

	advancedTechniques := `
	Advanced Techniques

- Subtests: Use t.Run for logically grouped test cases.
- Benchmarking: Measure performance with functions prefixed by Benchmark.
- Coverage: Use go test -cover to see how much of your code is covered by tests.
- Parallel Testing: Use t.Parallel() to run tests concurrently.	
	`
	fmt.Println(advancedTechniques)

	summary := `
	Summary

- Unit testing is a critical skill for ensuring code correctness.
- Use the testing package for writing simple and advanced tests.
- Employ table-driven tests for multiple test cases.
- Use mocking tools like testify to test code with external dependencies.
- Follow best practices for clear, maintainable tests.
	`
	fmt.Println(summary)
}
