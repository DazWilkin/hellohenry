package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("[handler] Entered")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "NOT SPECIFIED"
	}
	fmt.Fprintf(w, "Hello Henry: %s\n", target)
	log.Print("[handler] Exited")
}
func main() {
	log.Print("[main] Entered")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Print("[main] Exited")
}
