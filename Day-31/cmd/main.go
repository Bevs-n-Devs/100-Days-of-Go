package main

import (
	"fmt"
	"log"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// define the "User" struct - this will be the model for 'users' databse table
type User struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	Email   string `gorm:"size:255;unique"`
	Paaword string `gorm:"size:255"`
}

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nIntroduction to ORM (Object Relational Mapping) in Go")

	intro := `
ORM stands for Object-Relational Mapping, which is a programming technique that allows you to interact with a 
relational database (such as MySQL, PostgreSQL, SQLite) using object-oriented code instead of raw SQL queries. 
ORM tools, like GORM in Go, allow you to treat database tables as Go structs, making it easier to perform CRUD operations.	
`
	benefits := `
Benefits of ORM:
1. Abstraction: ORM abstracts the database operations, which can make your code easier to maintain and more readable.
2. Security: ORM reduces the risk of SQL injection attacks by automatically sanitizing queries.
3. Efficiency: It can reduce the amount of repetitive code, as it generates SQL queries based on struct definitions.
4. Portability: ORM frameworks like GORM allow you to switch between different relational databases easily.
`
	getStarted := `
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite  # or mysql, postgres driver
`
	bestPractices := `
Best Practices:
1. Use Structs to Define Database Tables: Define your database tables as Go structs, which makes it easier to work with.
2. Use Struct Tags for Column Names: Use struct tags to map database column names to struct fields.
3. Avoid Raw SQL: Use ORM tools to perform database operations instead of raw SQL queries.
4. Avoid Manual SQL Queries: Use ORM tools to generate SQL queries based on struct definitions.
5. Use Transactional Queries: Use transactions to ensure data consistency and avoid partial updates.
6. Avoid AutoMigrate in Production: While AutoMigrate is great for development, it can cause issues in production. Use proper migrations in production environments to manage schema changes.
`
	fmt.Println(intro)
	fmt.Println(benefits)
	fmt.Println(getStarted)
	fmt.Println(bestPractices)

	// initialise the database connection
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// migrate the database schema to create the "users" table
	db.AutoMigrate(&User{})

	// create a new user record
	user := User{Name: "John Doe", Email: "X0lZD@example.com", Paaword: "password123"}
	result := db.Create(&user) // insert into the DB

	if result.Error != nil {
		log.Fatal("Failed to create user:", result.Error)
	}
	fmt.Println("New user created:", user)

	// read user by ID
	var getUser User
	db.First(&getUser, 1) // select * from users where id = 1
	// error handling if user not found
	if getUser.ID == 0 {
		log.Fatal("User not found")
	}
	fmt.Println("User found:", getUser)

	// update user details
	getUser.Name = "Jane Doe"
	db.Save(&getUser) // update user details in the DB
	fmt.Println("User updated:", getUser)

	// delete user
	db.Delete(&User{}, 1) // delete user with ID 1
	// error handling
	if getUser.ID == 0 {
		log.Fatal("User not found")
	}
	fmt.Println("User deleted")

}
