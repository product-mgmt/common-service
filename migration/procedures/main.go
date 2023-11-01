package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// load config
	err := godotenv.Load(".env")
	if err != nil {
		msg := fmt.Sprintf("app.env loading error: %v", err)
		fmt.Println(msg)
		return
	}
	fmt.Println("config loaded!!")

	// Connect to the MySQL database
	conurl := os.Getenv("MYSQLDB_URL")
	dbname := os.Getenv("DATABASE_NAME")
	dburl := conurl + "/" + dbname + "?multiStatements=true"

	// Connect to the database.
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	fmt.Println("mysql db connected!!")

	// Specify the directory containing stored procedure files
	// procedureurl := os.Getenv("PROCEDURES_URL")

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	path, err := filepath.Abs(cwd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(path)

	// Read and execute each stored procedure file
	err = filepath.Walk(path+"/migration/procedures/migrate", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			sqlBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Execute the SQL statement
			_, err = db.Exec(string(sqlBytes))
			if err != nil {
				log.Printf("Error executing %s: %v\n", path, err)
			} else {
				log.Printf("Executed %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
