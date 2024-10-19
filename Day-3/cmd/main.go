package main

import "fmt"

// more on if statements: https://go.dev/doc/effective_go#if
func ifStatements() {
	fmt.Println("If statements: check age")
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	fmt.Println("If statements: check score")
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
}

// there is no difference between If statements vs Switch statements for performance: https://daryl-ng.medium.com/switch-vs-if-else-in-go-74b7d40aa70a
func switchStatements() {
	fmt.Println("Switch statements: check day")
	day := "Saturday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Unknown day.")
	}
}

// more on for loops: https://go.dev/doc/effective_go#for
func simpleForLoop() {
	fmt.Println("For Loop 1")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func simpleForLoop2() {
	fmt.Println("For Loop 2")
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func rangeLoop() {
	fmt.Println("For Loop 3: Using Range - printing the index with values")
	numbers := []int{10, 20, 30, 40}

	for i, num := range numbers {
		fmt.Printf("\nIndex: %d, Values: %d", i, num) // %d is used for printing integers in their decimals (base 10) format
	}

	fmt.Println("For Loop 3: Using Range - printing values without the index")
	// if the index is not needed we use _
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func controlStructures() {
	fmt.Println("Combining Control Structures to Create More Complex Loops")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("If statements, Switch statements & For loops")

	ifStatements()

	switchStatements()

	simpleForLoop()

	simpleForLoop2()

	rangeLoop()

	controlStructures()

	numberGuessingExercise()

}
