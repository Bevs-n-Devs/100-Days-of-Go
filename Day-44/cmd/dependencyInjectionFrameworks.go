package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Repository2 struct{}

func NewRepository2() *Repository2 {
	return &Repository2{}
}

type Service4 struct {
	repo *Repository2
}

func NewService4(repo *Repository2) *Service4 {
	return &Service4{repo: repo}
}

func (s *Service4) PerfomTask4() {
	fmt.Println("Tassk 4 performed with repository 2:", s.repo)
}

func dependencyInjectionFrameworks() {
	fmt.Println("\nDependency Injection Frameworks")

	intro := `
While Go's simplicity often makes DI frameworks unnecessary, they can be helpful in complex applications.

fx is a popular DI framework for Go. 
It provides a simple way to manage dependencies and is well-suited for small to medium-sized projects.

Benefits of fx:

	1. Automatically manages dependencies for you.
	2. Simplifies complex dependency trees.
	`
	fmt.Println(intro)

	app := fx.New(
		fx.Provide(NewRepository2),
		fx.Provide(NewService4),
		fx.Invoke(func(service *Service4) {
			service.PerfomTask4()
		}),
	)
	app.Run()
}
