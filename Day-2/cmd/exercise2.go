package main

import "fmt"

func RemainingOvenTime(actual int) int {
	const ovenTime = 40
	var result int = ovenTime - actual
	return result
}

func PreparationTime(numberOfLayers int) int {
	var result int = numberOfLayers * 2
	return result
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var result int = (numberOfLayers * 2) + actualMinutesInOven
	return result
}

func exercise2() {
	// 1. Define the expected oven time in minutes
	const ovenTime = 40
	fmt.Printf("\nOven time: %v mins\n", ovenTime)

	// 2. Calculate the remaining oven time in minutes
	var timeLeftInOven int = RemainingOvenTime(30)
	fmt.Printf("\nRemaining oven time: %v mins", timeLeftInOven)

	// 3. Calculate the preparation time in minutes
	var prepTime int = PreparationTime(2)
	fmt.Printf("\nPreparation time: %v mins", prepTime)

	// 4. Calculate the elapsed working time in minutes
	var timeElapsed int = ElapsedTime(3, 20)
	fmt.Printf("\nElapsed time: %v mins", timeElapsed)
}
