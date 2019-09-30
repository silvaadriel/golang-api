package main

import (
	"net/http"
)

//Router sets application endpoints
func Router() {
	http.HandleFunc("/client/", HandleClient)
}
