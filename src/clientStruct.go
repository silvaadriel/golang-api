package main

// Client represents a database entity
type Client struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	BirthDate string `json:"birthDate"`
}
