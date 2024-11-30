package main

import "fmt"

type Service1 struct{}

func (s *Service1) PerformTask() {
	fmt.Println("Performing task 1...")
}

type App1 struct {
	service *Service1
}

func NewApp1() *App1 {
	// App1 creates its own dependencies (tight coupling)
	return &App1{service: &Service1{}}
}

func (a *App1) Run() {
	a.service.PerformTask()
}

func exampleWithoutDI() {
	fmt.Println("\nExample Without Dependency Injection")

	app := NewApp1()
	app.Run()
}

/*
Problems:
	1. The App is tightly coupled to Service.
	2. You cannot replace Service with a mock or alternative implementation for testing.
*/
