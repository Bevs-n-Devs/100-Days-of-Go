package main

import "fmt"

func Factor1() {
	var factor1 = `
	1. Codebase: One Codebase, Tracked in Version Control
		- Principle: Each app should have a single codebase per repository, used across all environments (developement, staging, production).
		- Why: Ensure consistency and avoids environment-specific divergence.

		Example:
			- Use Git or similar for version control.
			- Avoid copying code manually between environments.

		Code Example:
			git clone https://github.com/example/your-app.git
	`
	fmt.Println(factor1)
}
