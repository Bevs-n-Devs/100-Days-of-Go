package main

import (
	"fmt"
	"sync"
)

/*
In this example the customer has a balance of £1000.
While they are adding £500 to their account, they are also being
charged £700 for a bill (ie. rent, mortgage etc).

Without a Mutex operation the outgoing £700 overrides the incoming £500,
which means the customer has lost £500 due to an error in the bank's system.

Using a mutex allows for one operation to be made, and then the other - meaning the customer
has the correct bank balance of £800 after the deposit has been made
and the bill has been withdrawn.
*/

var (
	mutex2  sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex2.Lock()
	fmt.Printf("Depositing £%d to account with balance £%d\n", value, balance)
	balance += value
	mutex2.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex2.Lock()
	fmt.Printf("Withdrawing £%d from account with balance £%d\n", value, balance)
	balance -= value
	mutex2.Unlock()
	wg.Done()
}

func bankBalanceMutexExample() {
	fmt.Println("\nBank Balance Mutex Example")

	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)

	// create the 2 Goroutines
	go withdraw(700, &wg)
	go deposit(500, &wg)

	// block until ALL WaitGroups have been completed
	wg.Wait()

	fmt.Printf("New balance £%d", balance)
}
