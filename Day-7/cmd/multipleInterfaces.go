package main

import "fmt"

/*
A type can have multiple interfaces.
This allows for even more flexibility.
*/

// Speaker2 interface
type Speaker2 interface {
	Speak4() string
}

// Mover interface
type Mover interface {
	Move() string
}

// Robot struct implements both Speaker and Mover
type Robot struct {
	id int
}

func (r Robot) Speak4() string {
	return "I am a robot"
}

func (r Robot) Move() string {
	return "I am moving"
}

func multipleInterfaces() {
	fmt.Println("\nUsing Multiple Interfaces")

	robot := Robot{id: 101}

	var speak Speaker2 = robot
	var move Mover = robot

	fmt.Println(speak.Speak4()) // using the Speaker2 interface
	fmt.Println(move.Move())    // using the Mover interface
}
