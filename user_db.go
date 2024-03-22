 
// user_db.go

package main

import (
    "database/sql"
    "log"
)

// InsertUser inserts a new user into the database
func InsertUser(db *sql.DB, user *User) error {
    _, err := db.Exec("INSERT INTO users (id, name, mobile, latitude) VALUES (?, ?, ?, ?)",
        user.ID, user.Name, user.Mobile, user.Latitude)
    if err != nil {
        log.Println("Error inserting user:", err)
        return err
    }
    return nil
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(db *sql.DB, id int) (*User, error) {
    var user User
    row := db.QueryRow("SELECT id, name, mobile, latitude FROM users WHERE id = ?", id)
    err := row.Scan(&user.ID, &user.Name, &user.Mobile, &user.Latitude)
    if err != nil {
        log.Println("Error retrieving user:", err)
        return nil, err
    }
    return &user, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(db *sql.DB, user *User) error {
    _, err := db.Exec("UPDATE users SET name = ?, mobile = ?, latitude = ? WHERE id = ?",
        user.Name, user.Mobile, user.Latitude, user.ID)
    if err != nil {
        log.Println("Error updating user:", err)
        return err
    }
    return nil
}

// DeleteUser deletes a user from the database
func DeleteUser(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        log.Println("Error deleting user:", err)
        return err
    }
    return nil
}
