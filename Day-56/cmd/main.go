package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello world, hello Yaw!")
	fmt.Println("\nThe Twelve-Factor App Approach")

	var definition = `
The Twelve-Factor App methodology provides best practices for building scalable, maintainable,
and portable software-as-a-service (SaaS) applications. It was introduced by Heroku and is widely
adpoted for modern application development.	
	`
	fmt.Println(definition)

	var whytwelvefactor = `
	Why Twelve-Factor App Matters
		1. Scalability: Facilitates horizontal scaling across distributed systems.
		2. Portability: Ensures applications can run in different environments (local, staging, production).
		3. Maintainability: Simplifies deployment, debugging, and adding features.
		4. DevOps Alignment: Encourages automation and smooth collaboration between developers and operations teams.
	`
	fmt.Println(whytwelvefactor)

	fmt.Println("\nThe 12 Factors")

	Factor1()
	Factor2()
	Factor3()
	Factor4()
	Factor5()
	Factor6()
	Factor7()
	Factor8()
	Factor9()
	Factor10()
	Factor11()
	Factor12()

	var benefits = `
	Benefits of the Twelve-Factor App Approach
		1. Scalability: Easily scale horizontally.
		2. Portability: Deploy across different platforms without modification.
		3. Resilience: Build fault-tolerant systems with minimal downtime.
		4. Speed: Enable rapid development cycles.
	`
	fmt.Println(benefits)

	var applyTwelveFactor = `
	How to Start Applying the Twelve Factors
		1. Audit your existing application for compliance with these principles.
		2. Implement changes incrementally (e.g., move configs to environment variables, introduce process isolation).
		3. Use tools and practices like CI/CD, Docker, and cloud platforms to align with the Twelve-Factor methodology.
	`
	fmt.Println(applyTwelveFactor)

	var summary = `
By Adopting the Twelve-Factor App principles, you'll ensure your application id robust,
maintainable, and ready to handle modern demands.
	`
	fmt.Println(summary)
}
