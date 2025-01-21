package main

import "fmt"

func Factor12() {
	var factor12 = `
	12. Admin Process: Run Admin/Management Tasks as One-Off Processes
		- Principle: Run tasks like database migrations as separate one-off commands.
		- Why: Keeps operational tasks decoupled from the app’s lifecycle.

		Example:
			- Database migrations using a Go library like 'migrate'.

		Code Example:
			migrate -path ./migrations -database "$DATABASE_URL" up
	`
	fmt.Println(factor12)
}
