package main

import "net/http"

func enableCors(responseWriter *http.ResponseWriter) {
	(*responseWriter).Header().Set("Access-Control-Allow-Origin", "*")
}
