package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Reading file")
	content, err := os.ReadFile("content/content.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Cannot read file")
	} else {
		fmt.Fprintf(w, string(content))
	}
}

func main() {
	log.Println("Starting listener")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
