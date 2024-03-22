 
package main

import (
	"database/sql"
	"log"
	
)

// InsertProduct inserts a new product into the database
func InsertProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec("INSERT INTO products (id,user_id, name, description, images, price) VALUES (?, ?, ?, ?, ?, ?)",
		product.ID,product.UserID, product.Name, product.Description, product.Images, product.Price)
	if err != nil {
		log.Println("Error inserting product:", err)
		return err
	}
	return nil
}

// GetProductByID retrieves a product by ID from the database
func GetProductByID(db *sql.DB, id int) (*Product, error) {
	var product Product
	row := db.QueryRow("SELECT id, user_id, name, description, images, price FROM products WHERE id = ?", id)
	err := row.Scan(&product.ID, &product.UserID, &product.Name, &product.Description, &product.Images, &product.Price)
	if err != nil {
		log.Println("Error retrieving product:", err)
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec("UPDATE products SET name = ?, description = ?, images = ?, price = ? WHERE id = ?",
		product.Name, product.Description, product.Images, product.Price, product.ID)
	if err != nil {
		log.Println("Error updating product:", err)
		return err
	}
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting product:", err)
		return err
	}
	return nil
}
