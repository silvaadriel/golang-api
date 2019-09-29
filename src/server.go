package main

import (
	"log"
	"net/http"
)

func main() {
	Router()
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
