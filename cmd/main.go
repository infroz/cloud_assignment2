package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"repocheck"
	"repocheck/commits"
)

const url = "/repocheck/v1/"

func main() {
	fmt.Println("Assignment2")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", repocheck.HandlerNil)

	http.HandleFunc(url+"commits", commits.HandlerCommits)
	//http.HandleFunc(url+"languages", repocheck.HandlerNil)
	//http.HandleFunc(url+"issues", repocheck.HandlerNil)
	//http.HandleFunc(url+"status", repocheck.HandlerNil)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
