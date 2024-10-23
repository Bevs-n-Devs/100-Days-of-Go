package main

import "fmt"

// define a bank account struct
type BankAccount struct {
	holder  string
	balance float64
}

// method to deposit money
func (account *BankAccount) deposit(amount float64) {
	account.balance += amount
	fmt.Printf("Deposited %.2f into %s's account\n", amount, account.holder)
}

// method to withdraw money
func (account *BankAccount) withdraw(amount float64) {
	if amount > account.balance {
		fmt.Printf("Insufficient funds in %s's account\n", account.holder)
	} else {
		account.balance -= amount
		fmt.Printf("Withdrew %.2f from %s's account\n", amount, account.holder)
	}
}

// method to check the balance (value reciever as it doesn't modify the struct)
func (account BankAccount) checkBalance() float64 {
	return account.balance
}

func bankBalanceExercise() {
	fmt.Println("\nBank Balance Exercise")

	// craete a new bank account
	account := BankAccount{
		holder:  "John Doe",
		balance: 1000.0,
	}

	// perfom some operations
	account.deposit(500)
	fmt.Println("Balance:", account.checkBalance())
	account.withdraw(200)
	fmt.Println("Balance:", account.checkBalance())
}
