package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The first app")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("The first app started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
