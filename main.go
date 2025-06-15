package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weatherForecast", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Today is gonna be a great sunny day for a go service!")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}