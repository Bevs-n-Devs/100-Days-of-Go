package main

import "fmt"

type Database interface {
	GetData(id int) string
}

func FetchData(db Database, id int) string {
	fmt.Println("Fetching data from DB...")
	return db.GetData(id)
}
