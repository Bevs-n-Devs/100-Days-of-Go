package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1. Generate a random number between 1 and 100 (use the math/rand package).
2. Prompt the user to guess the number.
3. Use a for loop to let the user guess multiple times until they get the correct number.
4. After each guess, provide feedback on whether the guess was too high, too low, or correct.
5. Exit the loop when the user guesses correctly, and display how many guesses they took.
*/

func numberGuessingExercise() {
	fmt.Println("Number Guessing Exercise")

	// seed the random number generator
	//  return a different value each time you run the program, so the seed is different every time.
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate a random number between 1 - 100
	number := randomNumber.Intn(100) + 1

	var guess int
	attempts := 0

	// start a for loop for the guessing game
	// this is a while loop
	for {
		fmt.Print("Guess a number between 1 - 100: ")
		fmt.Scanln(&guess)

		attempts++

		// add a control structure to check if the guess is too high, too low or correct
		if guess >= number {
			fmt.Println("Your guess is too high, guess again!")
			continue
		} else if guess <= number {
			fmt.Println("Your guess is too low, guess again!")
			continue
		} else {
			fmt.Println("You have guessed correctly!")
			fmt.Printf("\nThe correct number is %d, Th enumber of attempts you made were: %v", number, attempts)
			break
		}

		// if correct, break out of the loop and display the number of attempts
	}

}
