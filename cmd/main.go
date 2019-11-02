package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"repocheck"
	"repocheck/commits"
	"repocheck/issues"
	"repocheck/languages"
	"repocheck/status"
)

const url = "/repocheck/v1/"

func main() {
	fmt.Println("Assignment2")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", repocheck.HandlerNil)
	http.HandleFunc(url+"commits", commits.Handler)
	http.HandleFunc(url+"languages", languages.Handler)
	http.HandleFunc(url+"issues", issues.Handler)
	http.HandleFunc(url+"status", status.Handler)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
