package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/warehouse_db")
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	query, err := os.ReadFile("test.sql")
	if err != nil {
		fmt.Printf("Error reading SQL file: %v\n", err)
		return
	}

	var records int
	var maxCreatedAt string

	err = db.QueryRow(string(query)).Scan(&records, &maxCreatedAt)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}

	fmt.Printf("Count: %d, Max Created At: %s\n", records, maxCreatedAt)
}
