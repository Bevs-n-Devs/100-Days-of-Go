package mathutils

import (
	"fmt"
	"testing"
)

/*
When creating a test file, the convention is to give the file name an index of '_test.go' at the end.
And for the test function, we start the function name with 'Test..' followed by the name of the function we are testing.

TestAdd is our test function, which checks if Add(2, 3) returns 5.
t.Errorf is used to report a failure if the result isn’t as expected.
*/

func TestAdd(t *testing.T) {
	fmt.Println("Basics of Unit Testing in Go")

	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

/*
Table-driven testing is a pattern that makes it easy to test multiple inputs and
outputs by defining a table of cases and iterating over them.
This approach is common in Go.

Each test case in the cases slice has two inputs (a and b) and an expected output.
The loop iterates over the cases, calling Add and checking if it returns the correct value for each set of inputs.
*/
func TestAdd2(t *testing.T) {
	fmt.Println("\nTesting with Multiple Cases: Table-Driven Testing")

	cases := []struct {
		num1, num2 int
		expected   int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{10, -5, 5},
	}

	for _, c := range cases {
		result := Add(c.num1, c.num2)
		if result != c.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", c.num1, c.num2, result, c.expected)
		}
	}
}

/*
Go’s t.Run allows you to run each test case as a subtest, which is helpful for better test output, especially in larger tests.

Each case has a name field that describes it.
t.Run runs each case as a subtest, showing the case name in the test output.
*/
func TestAdd3(t *testing.T) {
	fmt.Println("\nUsing t.Run for Subsets")

	cases := []struct {
		name       string
		num1, num2 int
		expected   int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Nagative numbers", -1, 1, 0},
		{"Zeros", 0, 0, 0},
		{"Mixed Numbers", 10, -5, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Add(c.num1, c.num2)
			if result != c.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", c.num1, c.num2, result, c.expected)
			}
		})
	}
}

/*
It’s important to test edge cases to ensure your functions handle unexpected or extreme inputs gracefully.

Example: Testing Edge Cases
Suppose we modify Add2 to return an error if the sum exceeds 100
*/
func TestAdd4(t *testing.T) {
	fmt.Println("\nTesting Edge Case and Error Scenarios")

	cases := []struct {
		name        string
		num1, num2  int
		expected    int
		expectError bool
	}{
		{"Normal case", 2, 3, 5, false},
		{"Edge case over 100", 50, 51, 0, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := Add2(c.num1, c.num2)
			if c.expectError && err == nil {
				t.Errorf("Expected an error but got none")
			}
			if !c.expectError && err != nil {
				t.Errorf("did not expect an error but got one: %v", err)
			}
			if result != c.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", c.num1, c.num2, result, c.expected)
			}
		})
	}
}

/*
Go also supports benchmarking with functions that start with Benchmark.
Benchmarks are useful for performance testing.

Run benchmarks using go test -bench ..
b.N is an automatically determined number of iterations for a meaningful benchmark result.
*/
func BenchmarkAdd(b *testing.B) {
	fmt.Println("\nBenchmarking Tests in Go")
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

/*
When working with complex test or benchmark setups,
you can create helper functions and mark them as test helpers with t.Helper().
This ensures that if the helper function fails, the line number reported will be where the helper was called, not inside the helper.
*/
func checkResult(t *testing.T, result, expected int) {
	t.Helper() // marks this function as helper
	if result != expected {
		t.Errorf("got %d; expected %d", result, expected)
	}
}

func TestAdd5(t *testing.T) {
	fmt.Println("Using testing.T and testing.B Helper functions")
	result := Add(2, 3)
	checkResult(t, result, 5)
}

/*
MORE ON TESTING

Test coverage shows the percentage of your code covered by tests, which is helpful for spotting untested code:

go test -cover

To generate a detailed coverage report:

go test -coverprofile=coverage.out
go tool cover -html=coverage.out

This generates an HTML report highlighting which parts of your code were tested.

*/

/*
BEST PRACTICES FOR TESTING IN GO

1. Use Table-Driven Tests for functions with multiple cases.
2. Keep Tests Independent: Tests should be isolated and not depend on each other.
3. Avoid Complex Logic in Tests: Aim for simple, clear assertions. (KISS!: Keep It Simple Stupid!)
4. Use t.Run for Subtests to provide clear output and structure.
5. Test Edge Cases to ensure code handles unexpected inputs.
6. Aim for 100% Coverage but focus on testing critical paths and edge cases.
*/
