package main

// define the struct for the quiz and prepopulate it with questions and answers
type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"` // multiple choice options
	Answer   string   `json:"answer"`  // correct answer
}

var questions = []Question{
	{
		ID:       1,
		Question: "What is the capital of France?",
		Options:  []string{"Paris", "London", "Berlin", "Rome"},
		Answer:   "Paris",
	},
	{
		ID:       2,
		Question: "What is the largest planet in our solar system?",
		Options:  []string{"Earth", "Mars", "Jupiter", "Saturn"},
		Answer:   "Jupiter",
	},
	{
		ID:       3,
		Question: "What is the most popular programming language?",
		Options:  []string{"Python", "Java", "C++", "JavaScript"},
		Answer:   "Python",
	},
	{
		ID:       4,
		Question: "What is the oldest programming language?",
		Options:  []string{"Python", "Java", "C++", "Assembly"},
		Answer:   "Assembly",
	},
	{
		ID:       5,
		Question: "What is the my favourite football team?",
		Options:  []string{"Manchester United", "Liverpool", "Chelsea", "Arsenal"},
		Answer:   "Arsenal",
	},
}
