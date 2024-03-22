

package main

// Product represents a product entity
type Product struct {
    ID          int      `json:"id"`
    UserID      int      `json:"user_id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Images      string   `json:"images"`
    Price       float64  `json:"price"`
}

type User struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Mobile   string  `json:"mobile"`
    Latitude float64 `json:"latitude"`
}