package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Product struct
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

// Environment variable keys for database credentials
const (
	dbUserKey     = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	dbNameKey     = "DB_NAME"
	dbHostKey     = "DB_HOST"
)

// Connect to the database
func connectDB() (*sql.DB, error) {
	dbUser := os.Getenv(dbUserKey)
	dbPassword := os.Getenv(dbPasswordKey)
	dbName := os.Getenv(dbNameKey)
	dbHost := os.Getenv(dbHostKey)

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	return sql.Open("postgres", connStr)
}

// Insert a product
func insertProduct(db *sql.DB, product Product) (int, error) {
	query := `INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, product.Name, product.Description, product.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insertProduct: %v", err)
	}
	return id, nil
}

// Get a product by ID
func getProductByID(db *sql.DB, id int) (Product, error) {
	var product Product
	query := `SELECT id, name, description, price FROM products WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return product, fmt.Errorf("getProductByID: %v", err)
	}
	return product, nil
}

// Update a product
func updateProduct(db *sql.DB, product Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4`
	_, err := db.Exec(query, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("updateProduct: %v", err)
	}
	return nil
}

// Delete a product
func deleteProduct(db *sql.DB, id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("deleteProduct: %v", err)
	}
	return nil
}

// Update a product and log the action using a transaction
func updateProductWithTransaction(db *sql.DB, product Product) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("start transaction: %v", err)
	}

	queryUpdate := `UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4`
	if _, err := tx.Exec(queryUpdate, product.Name, product.Description, product.Price, product.ID); err != nil {
		tx.Rollback()
		return fmt.Errorf("updateProductWithTransaction - update: %v", err)
	}

	queryLog := `INSERT INTO product_logs (product_id, action) VALUES ($1, $2)`
	if _, err := tx.Exec(queryLog, product.ID, "update"); err != nil {
		tx.Rollback()
		return fmt.Errorf("updateProductWithTransaction - log: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}
	return nil
}

func main() {
	fmt.Println("hello world, hello Yaw!")

	fmt.Println("\nSQL Queries with PostreSQL in Go")

	// Connect to the database
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Example: Insert a product
	newProduct := Product{Name: "Sample Product", Description: "A sample description", Price: 19.99}
	productID, err := insertProduct(db, newProduct)
	if err != nil {
		log.Printf("Error inserting product: %v", err)
	} else {
		log.Printf("Inserted product with ID %d", productID)
	}

	// Example: Get a product
	product, err := getProductByID(db, productID)
	if err != nil {
		log.Printf("Error retrieving product: %v", err)
	} else {
		log.Printf("Retrieved product: %+v", product)
	}

	// Example: Update a product with transaction
	product.Name = "Updated Product Name"
	product.Description = "Updated description"
	product.Price = 25.99
	err = updateProductWithTransaction(db, product)
	if err != nil {
		log.Printf("Error updating product with transaction: %v", err)
	} else {
		log.Println("Product updated with transaction")
	}

	// Example: Delete a product
	err = deleteProduct(db, productID)
	if err != nil {
		log.Printf("Error deleting product: %v", err)
	} else {
		log.Printf("Deleted product with ID %d", productID)
	}
}

/*
1. Understand and execute CRUD operations with SQL queries in Go.
2. Learn to manage transactions for reliable data manipulation.
3. Practice handling query errors and preparing statements.

Topics Covered:
- CRUD operations with SQL queries
- Handling SQL injections with prepared statements
- Working with transactions for data consistency
- Query error handling
*/
