package main

import "fmt"

func Example3() {
	var scenario = `
Example 3: Avoid Building for Hypothetical Scalability

	Scenario:
	You're designing a small blog application that allows users to post articles.

	Premature Implementation:
	You decide to use a distributed database and implement a sophisticated microservices architecture becasue you expect
	the app to scale to millions of users someday.

	Problems:
		- Premature Optimization: Your app may never reach that scale.
		- Increased Costs: Distributed systems and microservices are harder and more expensive to develop and maintain.
		- Slower Development: Adding unnecessary infrastructure delays launching your app.
	
	YAGNI Approach:
	Start with a simple monolithic architecture and a single database. If the app grows and you need to scale, you can
	migrate to a distributed system when necessary.
	`
	fmt.Println(scenario)
}
