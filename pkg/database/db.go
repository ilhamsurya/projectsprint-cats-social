package main

import (
	"database/sql"
	"fmt"
	"log"
	"projectsphere/cats-social/pkg/utils/config"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	DBConnect *sql.DB
	err       error
)

var (
	host     = config.Get().DB.Postgre.Host
	port     = config.Get().DB.Postgre.Port
	user     = config.Get().DB.Postgre.User
	password = config.Get().DB.Postgre.Pass
	dbName   = config.Get().DB.Postgre.Name
)

func Run() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	DBConnect, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = DBConnect.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set connection parameters
	DBConnect.SetMaxOpenConns(10)                 // Maximum number of open connections
	DBConnect.SetMaxIdleConns(5)                  // Maximum number of idle connections
	DBConnect.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	return DBConnect
}

func main() {
	db := Run()
	defer DBConnect.Close() // Close the database connection when done

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up() // Apply all available migrations
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Migrations applied successfully!")
}
