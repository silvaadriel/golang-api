package main

import "database/sql"

// DBConnect opens a postgres database connection
func DBConnect() (*sql.DB, error) {
	connStr := "host=postgres dbname=crud user=postgres password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	return db, err
}
