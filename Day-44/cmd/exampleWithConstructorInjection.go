package main

import "fmt"

type Repository struct {
	db string
}

func NewRepository(db string) *Repository {
	return &Repository{db: db}
}

type Service3 struct {
	repo *Repository
}

func NewService3(repo *Repository) *Service3 {
	return &Service3{repo: repo}
}

func (s *Service3) FetchData() {
	fmt.Printf("Fetching data from DB: %s\n", s.repo.db)
}

func exampleWithConstructorInjection() {
	fmt.Println("\nExample With Constructor Injection")

	// inject repository dependency into service
	repo := NewRepository("PostgreSQL")
	service := NewService3(repo)
	service.FetchData()
}
