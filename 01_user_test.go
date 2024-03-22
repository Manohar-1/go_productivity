package main 

import (
    "database/sql"
    "testing"
    "fmt"
)
// TestUserCRUDOperations tests user CRUD operations
 func TestUserCRUDOperations(t *testing.T) {
    // Open a test database connection
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database_product_management")
    if err != nil {
        t.Fatalf("Error opening database connection: %v", err)
    }
    defer db.Close()

    // Prepare test data
    user := &User{ 
        ID: 1,
        Name:     "Test User",
        Mobile:   "1234567890",
        Latitude: 40.7128,
    }
    user_2 := &User{ 
        ID: 169,
        Name:     "Test User2",
        Mobile:   "1234567890",
        Latitude: 40.7128,
    }
    user_3 := &User{ 
        ID:3,
        Name:     "Test User3",
        Mobile:   "1234567890",
        Latitude: 40.7128,
    }

    

    // Test InsertUser
    if err := InsertUser(db, user); err != nil {
        t.Fatalf("Insert User failed: %v", err)
    }

    fmt.Println("insertedUser1")
    // Test GetUserByID
    insertedUser_1, err := GetUserByID(db, user.ID)
    if err != nil {
        t.Fatalf("GetUserByID failed: %v", err)
    }
    fmt.Println(insertedUser_1)
    if err := InsertUser(db, user_2); err != nil {
        t.Fatalf("Insert User failed: %v", err)
    }
    
    fmt.Println("insertedUser2")

    // Test GetUserByID
    insertedUser_2, err := GetUserByID(db, user_2.ID)
    if err != nil {
        t.Fatalf("GetUserByID failed: %v", err)
    }
   
    fmt.Println(insertedUser_2)

    if err := InsertUser(db, user_3); err != nil {
        t.Fatalf("Insert User failed: %v", err)
    }

    fmt.Println("insertedUser3")

    // Assert retrieved user matches the expected user
    // (Compare fields of insertedUser with fields of user)

    // Test UpdateUser
    user.Name = "Updated User_126"
    if err := UpdateUser(db, user); err != nil {
        t.Fatalf("UpdateUser failed: %v", err)
    }

    // Test DeleteUser
    if err := DeleteUser(db, user_2.ID); err != nil {
        t.Fatalf("DeleteUser failed: %v", err)
    }
}