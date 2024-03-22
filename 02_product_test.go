package main 

import ( 
    "database/sql"
    "testing"
    "fmt"
)

func TestProductCRUDOperations(t *testing.T) {
    // Open a test database connection
    testDB, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database_product_management")
    if err != nil {
        t.Fatalf("Error opening database connection: %v", err)
    }
    defer testDB.Close()

    // Prepare test data
    product := &Product{
        ID: 1004,
        UserID:      1,
        Name:        "Test Product",
        Description: "This is a test product",
        Images:      "image1.jpg",
        Price:       9.99,
    }
    product_2 := &Product{ 
        ID:1005,
        UserID : 1,
        Name : "Test PRoduct 2",
        Description:"This is test product_2",
        Images: "image_2.jpg",
        Price:9.80,
    }
    product_3 := &Product{ 
        ID:1006,
        UserID : 1,
        Name : "Test PRoduct 3",
        Description:"This is test product_3",
        Images: "image_3.jpg",
        Price:9.86,
    }


    // Test InsertProduct
    if err := InsertProduct(testDB, product); err != nil {
        t.Fatalf("InsertProduct failed: %v", err)
    } 

    fmt.Println("inserted product1");

    if err := InsertProduct(testDB, product_2); err != nil {
        t.Fatalf("InsertProduct failed: %v", err)
    } 
    fmt.Println("inserted product2");

    if err := InsertProduct(testDB, product_3); err != nil {
        t.Fatalf("InsertProduct failed: %v", err)
    }
    fmt.Println("inserted product3");
    // Test GetProductByID
    insertedProduct_1, err := GetProductByID(testDB, product.ID)
    if err != nil {
        t.Fatalf("GetProductByID failed: %v", err)
    }
    
    fmt.Println(insertedProduct_1)
    // Assert retrieved product matches the expected product
    // (Compare fields of insertedProduct with fields of product)

    // Test UpdateProduct
    product_2.Name = "Updated Product_2"
    if err := UpdateProduct(testDB, product_2); err != nil {
        t.Fatalf("UpdateProduct failed: %v", err)
    }

    // Test DeleteProduct
    if err := DeleteProduct(testDB, product_3.ID); err != nil {
        t.Fatalf("DeleteProduct failed: %v", err)
    }
}
