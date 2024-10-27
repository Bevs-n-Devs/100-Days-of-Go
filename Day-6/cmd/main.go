package main

import "fmt"

// define struct Person
type Person struct {
	name   string
	age    int
	gender string
}

// a method is a function associated with a type (usaually a struct)
// You can think of methods as functions that "belong" to a specific struct.
func (p Person) greet() {
	fmt.Printf("Hello, my name is %s, and I am %d years old\n", p.name, p.age)
}

// Value receivers receive a copy of the struct, meaning changes made within the method do not affect the original struct.
func (p Person) valueRecievers() {
	p.age++ // this will not affect the original struct
}

// Pointer receivers receive a pointer to the struct, meaning changes made within the method will modify the original struct.
func (p *Person) pointerReciever() {
	p.age++ // this will affect the original struct
}

// Go doesn’t have traditional object-oriented inheritance, but you can achieve similar behavior through struct embedding.
// One struct can include another struct, which allows you to group related functionality.
type Address struct {
	city    string
	country string
}

type People struct {
	name    string
	age     int
	Address // Embedded struct
}

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nStructs and Methods")
	// creating a new struct instance
	person1 := Person{
		name:   "Simon",
		age:    39,
		gender: "male",
	}

	// assessing fields of the struct
	fmt.Println("Name:", person1.name)
	fmt.Println("Age:", person1.age)
	fmt.Println("Gender:", person1.gender)

	// you can also modify the fileds
	person1.name = "Yaw"
	person1.age = 38
	fmt.Println("Updated Name:", person1.name)
	fmt.Println("Updated Age:", person1.age)

	// initiating structs can be done in THREE ways
	// one like about and the other two like below
	fmt.Println("\nWays to Initiate a Struct")
	person2 := Person{"Mary", 29, "female"}
	fmt.Println("Create structs without field names:", person2)

	// can create structs using the zero value - by not specifying any values this defaults to 0
	var person3 Person
	fmt.Println("Zero value struct:", person3)

	fmt.Println("\nMethods in Go")
	// call the struct method
	person1.greet()

	fmt.Println("\nValue Recievers vs Value Pointers")
	person4 := Person{"Alice", 25, "female"}
	person4.valueRecievers()
	fmt.Println("Age after incrementing:", person4.age) // 25

	person4.pointerReciever()
	fmt.Println("Age after incrementing:", person4.age)

	bankBalanceExercise()

	// anonymous structs - useful when for small temporary task and don't want to define a separate named type
	// they are good for quick throwaway data structures that don't need to be reused throughout the program
	fmt.Println("\nAnonymous Structs")
	person := struct {
		name string
		age  int
	}{
		name: "John",
		age:  34,
	}

	// access the fields
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)

	fmt.Println("\nComposition and Emdedding Structs")
	randomPerson := People{
		name: "Max",
		age:  40,
		Address: Address{
			city:    "London",
			country: "UK",
		},
	}

	// access field of both stuct and embedded struct
	fmt.Println("Name:", randomPerson.name)
	fmt.Println("Age:", randomPerson.age)
	fmt.Println("City:", randomPerson.city)
	fmt.Println("Country:", randomPerson.country)

	twoSumAlgorithm()
}

// When to Use Value vs Pointer Receivers?
// Value receivers are fine if the method doesn’t need to modify the struct, or if the struct is small.
// Pointer receivers should be used if the method needs to modify the struct or if the struct is large (to avoid copying large amounts of data).
