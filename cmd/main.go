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
	"repocheck/webhooks"
	"time"
)

const url = "/repocheck/v1/"

func init() {
	status.StartTime = time.Now()
}

func main() {
	fmt.Println("Assignment2")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := webhooks.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer webhooks.Close()

	http.HandleFunc("/", repocheck.HandlerNil)
	http.HandleFunc(url+"commits", commits.Handler)
	http.HandleFunc(url+"languages", languages.Handler)
	http.HandleFunc(url+"issues", issues.Handler)
	http.HandleFunc(url+"status", status.Handler)
	http.HandleFunc(url+"webhooks", webhooks.WebhookHandler)
	http.HandleFunc(url+"webhook/", webhooks.WebhookHandlerID)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
