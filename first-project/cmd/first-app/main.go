package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the first App")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("first App started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
