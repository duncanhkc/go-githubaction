package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The third app")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("The third app started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
