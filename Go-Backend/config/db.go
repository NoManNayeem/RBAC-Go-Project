package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() error {
	// Define connection parameters
	server := os.Getenv("DB_SERVER")     // e.g., "localhost"
	port := os.Getenv("DB_PORT")         // e.g., "1433"
	user := os.Getenv("DB_USER")         // e.g., "sa"
	password := os.Getenv("DB_PASSWORD") // Your password
	database := os.Getenv("DB_NAME")     // e.g., "RBAC"

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, server, port, database)

	// Open a connection to the database
	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Connected to MS SQL Server successfully!")
	return nil
}
