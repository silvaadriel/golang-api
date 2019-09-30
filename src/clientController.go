package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// HandleClient parses request and delegates for proper function
func HandleClient(responseWriter http.ResponseWriter, request *http.Request) {
	stringID := strings.TrimPrefix(request.URL.Path, "/client/")
	ID, _ := strconv.Atoi(stringID)

	switch {
	case request.Method == "GET" && ID > 0:
		show(responseWriter, request, ID)
	case request.Method == "GET":
		index(responseWriter, request)
	case request.Method == "POST":
		store(responseWriter, request)
	case request.Method == "PUT":
		update(responseWriter, request, ID)
	case request.Method == "DELETE":
		destroy(responseWriter, request, ID)
	default:
		responseWriter.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(responseWriter, "Sorry... :(")
	}
}

func show(responseWriter http.ResponseWriter, request *http.Request, ID int) {
	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var client Client

	sql := "select * from clients where id = $1"
	db.QueryRow(sql, ID).Scan(&client.ID, &client.Name, &client.LastName,
		&client.Email, &client.BirthDate)

	json, _ := json.Marshal(client)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprint(responseWriter, string(json))
}

func index(responseWriter http.ResponseWriter, request *http.Request) {
	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select * from clients"
	rows, _ := db.Query(sql)
	defer rows.Close()

	var clients []Client

	for rows.Next() {
		var client Client
		rows.Scan(&client.ID, &client.Name, &client.LastName, &client.Email,
			&client.BirthDate)
		clients = append(clients, client)
	}

	json, _ := json.Marshal(clients)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprint(responseWriter, string(json))
}

func store(responseWriter http.ResponseWriter, request *http.Request) {
	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
	responseWriter.WriteHeader(http.StatusOK)

	var client Client
	json.Unmarshal([]byte(body), &client)

	sql := `INSERT INTO clients(name, last_name, email, birth_date)
	VALUES($1, $2, $3, $4) RETURNING id`
	db.QueryRow(sql, client.Name, client.LastName,
		client.Email, client.BirthDate).Scan(&client.ID)

	json, _ := json.Marshal(client)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprint(responseWriter, string(json))
}

func update(responseWriter http.ResponseWriter, request *http.Request, ID int) {
	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
	responseWriter.WriteHeader(http.StatusOK)

	var client Client
	json.Unmarshal([]byte(body), &client)

	sql := `UPDATE clients SET name = $1, last_name = $2, email = $3, birth_date = $4
	WHERE id = $5 RETURNING id`
	db.QueryRow(sql, client.Name, client.LastName,
		client.Email, client.BirthDate, ID).Scan(&client.ID)

	json, _ := json.Marshal(client)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprint(responseWriter, string(json))
}

func destroy(responseWriter http.ResponseWriter, request *http.Request, ID int) {
	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var client Client

	sql := "DELETE FROM clients WHERE id = $1 RETURNING *"
	db.QueryRow(sql, ID).Scan(&client.ID, &client.Name, &client.LastName,
		&client.Email, &client.BirthDate)

	json, _ := json.Marshal(client)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprint(responseWriter, string(json))
}
