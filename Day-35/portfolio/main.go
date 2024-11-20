package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nMy Portfolio Website in Go")

	// create a new Gin router
	r := gin.Default()

	// serve static files and templates
	r.Static("/static", "./static") // serve CSS or other static files
	r.LoadHTMLGlob("templates/*")   // serve HTML templates

	// define routes related to the portfolio
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// handle post request from /quiz form
	r.POST("/quiz", func(ctx *gin.Context) {
		// get the form data
		question1 := ctx.PostForm("question1")
		question2 := ctx.PostForm("question2")
		question3 := ctx.PostForm("question3")
		question4 := ctx.PostForm("question4")
		question5 := ctx.PostForm("question5")

		// initialize the answers
		var question1Answer string
		var question2Answer string
		var question3Answer string
		var question4Answer string
		var question5Answer string

		// validate the answers
		if question1 != "Python" {
			question1Answer = "WRONG ANSWER! Correct Answer: Python"
		} else {
			question1Answer = "CORRECT ANSWER!"
		}

		if question2 != "Blue" {
			question2Answer = "WRONG ANSWER! Correct Answer: Blue"
		} else {
			question2Answer = "CORRECT ANSWER!"
		}

		if question3 != "Fufu" {
			question3Answer = "WRONG ANSWER! Correct Answer: Fufu, best served with peanut soup!"
		} else {
			question3Answer = "CORRECT ANSWER! Fufu, best served with peanut soup!"
		}

		if question4 != "2017" {
			question4Answer = "WRONG ANSWER! Correct Answer: 2017. I built my first application in 6 weeks"
		} else {
			question4Answer = "CORRECT ANSWER! I built my first application in 6 weeks"
		}

		if question5 != "2022" {
			question5Answer = "WRONG ANSWER! Correct Answer: 2022. My first job was with Amadeus IT."
		} else {
			question5Answer = "CORRECT ANSWER! My first job was with Amadeus IT."
		}

		// render the results page
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			// we cannot use question1Answer == "CORRECT ANSWER!" directly so instead
			// we create a map and pass it to the template
			"checkAnswer1":    question1Answer == "CORRECT ANSWER!",
			"question1Answer": question1Answer,
			"checkAnswer2":    question2Answer == "CORRECT ANSWER!",
			"question2Answer": question2Answer,
			"checkAnswer3":    question3Answer == "CORRECT ANSWER! Fufu, best served with peanut soup!",
			"question3Answer": question3Answer,
			"checkAnswer4":    question4Answer == "CORRECT ANSWER! I built my first application in 6 weeks",
			"question4Answer": question4Answer,
			"checkAnswer5":    question5Answer == "CORRECT ANSWER! My first job was with Amadeus IT.",
			"question5Answer": question5Answer,
		})
	})

	// start the server
	r.Run(":8080")
}
