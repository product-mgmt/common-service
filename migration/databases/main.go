package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// load config
	err := godotenv.Load(".env")
	if err != nil {
		msg := fmt.Sprintf("app.env loading error: %v", err)
		fmt.Println(msg)
		os.Exit(1)
	}

	// Replace with your MySQL connection details
	conurl := os.Getenv("MYSQLDB_URL")
	dbname := os.Getenv("DATABASE_NAME")
	migrationurl := os.Getenv("DB_MIGRATION_URL")

	// Define your MySQL connection parameters
	db, err := sql.Open("mysql", conurl+"/")
	if err != nil {
		log.Fatal(err)
	}

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dburl := conurl + "/" + dbname

	dbURL := "mysql://" + dburl

	// Create a new migration instance
	m, err := migrate.New(migrationurl, dbURL)
	if err != nil {
		fmt.Println("Error creating migration instance:", err)
		os.Exit(1)
	}
	defer m.Close()

	args := os.Args[1]
	if args == "up" {
		// Apply migrations
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Error applying migrations:", err)
			os.Exit(1)
		}
	}

	if args == "down" {
		// Apply migrations
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Error applying migrations:", err)
			os.Exit(1)
		}

		_, err = db.Exec("DROP DATABASE IF EXISTS " + dbname)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Migrations applied successfully!")
}
