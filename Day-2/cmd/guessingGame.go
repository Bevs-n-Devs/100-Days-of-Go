package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
This program generates a random number between 1 and 100,
and the user has to guess the number. After each guess,
the program provides feedback on whether the guess is
too high, too low, or correct.
*/

func guessingGame() {
	fmt.Println("\n", strings.Repeat("*", 4), "Guessing Game with Go!", strings.Repeat("*", 4))

	fmt.Println("INSTRUCTIONS:")
	fmt.Println("Guess a number between 1 - 100 and the program will tell you if you are: TOO HIGH, TOO LOW or CORRECT!")

	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	target := rand.Intn(100) + 1     // generate a random number between 1 - 100
	var guess int

	fmt.Println("\nWelcome to the guessing game!")
	fmt.Println("The program has chosen a numnber between 1 and 100. Can you guess it? You have 5 attempts.")

	for attempts := 1; attempts <= 5; attempts++ {
		fmt.Println("Enter your guess: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("TOO LOW!")
		} else if guess > target {
			fmt.Println("TOO HIGH!")
		} else {
			fmt.Printf("CONGRATULATIONS! You guessed the number if the %d attempts", attempts)
			break
		}
	}
}
