package main

import "fmt"

type Service2 interface {
	PerformTask2()
}

type RealService struct{}

func (s *RealService) PerformTask2() {
	fmt.Println("Performing task 2 with RealService...")
}

type MockService struct{}

func (s *MockService) PerformTask2() {
	fmt.Println("Performing task 2 with MockService...")
}

type App2 struct {
	service Service2
}

func NewApp2(service Service2) *App2 {
	// injecting the dependency
	return &App2{service: service}
}

func (a *App2) RunApp2() {
	a.service.PerformTask2()
}

func exampleWithDI() {
	fmt.Println("\nExample With Dependency Injection")
	// use RealService in production
	realService := &RealService{}
	app := NewApp2(realService)
	app.RunApp2()

	// use the MockService in test
	mockService := &MockService{}
	testApp := NewApp2(mockService)
	testApp.RunApp2()
}

/*
Benefits:
	1. The App is loosely coupled to Service.
	2. You can replace Service with a mock or alternative implementation for testing.
*/
