package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showQuizPage(c *gin.Context) {
	// pass questions to the template
	c.HTML(http.StatusOK, "quiz.html", questions)
}

func submitQuiz(c *gin.Context) {
	// collect answers from the HTML form
	userAnswers := make(map[int]string)
	for _, question := range questions {
		userAnswer := c.PostForm(string(question.ID))
		userAnswers[question.ID] = userAnswer
	}

	// validate the answers
	results := []map[string]string{}
	for _, question := range questions {
		correct := userAnswers[question.ID] == question.Answer
		results = append(results, map[string]string{
			"Question":   question.Question,
			"yourAnswer": userAnswers[question.ID],
			"correct":    question.Answer,
			"result":     ifThenElse(correct, "correct", "incorrect"),
		})
	}

	// render the results page
	c.HTML(http.StatusOK, "results.html", results)
}

// utility function for conditional evaluations
func ifThenElse(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}

// define the routes, quiz logic, and the main function
func main() {
	r := gin.Default()

	// serve the static files and templates
	r.Static("/static", "./static") // serve CSS or other static files
	r.LoadHTMLGlob("templates/*")   // serve HTML templates

	// define routes related to the portfolio
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/projects", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "projects.html", nil)
	})

	// quiz routes
	r.GET("/quiz", showQuizPage)
	r.POST("/quiz", submitQuiz)

	// start the server
	r.Run(":8080")
}
