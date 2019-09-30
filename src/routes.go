package main

import (
	"net/http"
)

//Router
func Router() {
	http.HandleFunc("/client", HandleClient)
}
