package main

import "fmt"

func Factor3() {
	var factor3 = `
	3. Config: Store Configuration in the Environment
		- Pinciple: Store app-sepcific configuration (eg. databasae URLs, API keys, variables, not in code).
		- Why: Allows changes without to modifying 

		Code Example:
			export DATABASE_URL="postgres://username:password@localhost:5432/db"

	`
	fmt.Println(factor3)
}
