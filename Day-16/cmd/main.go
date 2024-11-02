package main

import "fmt"

/*
1. Set up an HTTP server in.
2. Define basic REST API routes for a resource (e.g., Task).
3. Implement CRUD (Create, Read, Update, Delete) operations for the resource.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nBuilding A CRUD App With Go")

	server()
}

/*
TESTING YOUR API

- Create a new task:
curl -X DELETE http://localhost:8080/tasks/1

- Get all tasks:
curl http://localhost:8080/tasks

- Update a task:
curl -X PUT -d '{"name": "Learn Go Concurrency", "completed": true}' -H "Content-Type: application/json" http://localhost:8080/tasks/1

- Delete a task:
curl -X DELETE http://localhost:8080/tasks/1

*/
