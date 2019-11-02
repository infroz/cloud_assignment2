package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
	Handles webhook registration (POST) and lookup (GET) requests.
	Expects WebhookRegistration struct body in request.
*/
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Expects incoming body in terms of WebhookRegistration struct

		err := json.NewDecoder(r.Body).Decode(&Wh)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
		}

		Wh.Time = time.Now()

		Wh.Event = strings.ToLower(Wh.Event)

		err = Add()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Webhook " + Wh.Url + " has been registered.")
	case http.MethodGet:
		var webhook []Webhook // Webhook db
		webhook, err := Read()
		if err != nil {
			log.Fatalln(err)
		}
		// For now just return all webhooks, don't respond to specific resource requests
		err = json.NewEncoder(w).Encode(webhook)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		err := json.NewDecoder(r.Body).Decode(&Wh)
		if err != nil {
			log.Fatalln(err)
		}

		err = Delete(Wh.ID)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		http.Error(w, "Invalid method "+r.Method, http.StatusBadRequest)
	}
}

func WebhookHandlerID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var webhook []Webhook

		webhook, err := Read()
		if err != nil {
			log.Fatalln(err)
		}
		parts := strings.Split(r.URL.Path, "/")

		for i := range webhook {
			if webhook[i].ID == parts[4] {
				err = json.NewEncoder(w).Encode(webhook[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	default:
		http.Error(w, "Invalid method "+r.Method, http.StatusBadRequest)
	}
}

func URLCaller(event string, params string, timeStamp time.Time) {
	var webhook []Webhook
	webhook, err := Read()
	if err != nil {
		log.Fatalln(err)
	}

	for _, i := range webhook {
		if i.Event == event {
			var request = Invocation{Event: event, Parameters: params, Timestamp: timeStamp.String()}

			requestBody, err := json.Marshal(request)
			if err != nil {
				log.Fatalln(err)
			}

			resp, err := http.Post(i.Url, "json", bytes.NewReader(requestBody))
			if err != nil {
				log.Fatalln(err)
			}

			response, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("Webhook invoked, body: " + string(response))
		}
	}

}
