package concurrencypatterns

import (
	"fmt"
)

/*
In this pattern, the output of one stage is the input for the next. It’s commonly used for transforming data through multiple steps.
*/

// generator stage
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// square stage
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func PipelinePattern() {
	fmt.Println("\nConcurrency Patterns: Pipeline Patterns")

	nums := generator(1, 2, 3, 4)
	results := square(nums)

	for result := range results {
		fmt.Println(result)
	}
}
