package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type WebhookRegistration struct {
	Url   string `json:"url"`
	Event string `json:"event"`
}

var webhooks []WebhookRegistration // Webhook db

/*
	Handles webhook registration (POST) and lookup (GET) requests.
	Expects WebhookRegistration struct body in request.
*/
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Expects incoming body in terms of WebhookRegistration struct
		webhook := WebhookRegistration{}
		err := json.NewDecoder(r.Body).Decode(&webhook)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
		}
		webhooks = append(webhooks, webhook)
		// Note: Approach does not guarantee persistence or permanence of resource id (for CRUD)
		fmt.Fprintln(w, len(webhooks)-1)
		fmt.Println("Webhook " + webhook.Url + " has been registered.")
	case http.MethodGet:
		// For now just return all webhooks, don't respond to specific resource requests
		err := json.NewEncoder(w).Encode(webhooks)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Invalid method "+r.Method, http.StatusBadRequest)
	}
}

/*
*  Invokes the web service to trigger event. Currently only responds to POST requests.
 */
func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Println("Received POST request...")
		for _, v := range webhooks {
			go CallUrl(v.Url, "Response on registered event in webhook demo: "+v.Event)
		}
	default:
		http.Error(w, "Invalid method "+r.Method, http.StatusBadRequest)
	}
}

/*
	Calls given URL with given content and awaits response (status and body).
*/
func CallUrl(url string, content string) {
	fmt.Println("Attempting invocation of url " + url + " ...")
	res, err := http.Post(url, "string", bytes.NewReader([]byte(content)))
	if err != nil {
		fmt.Println("Error in HTTP request: " + err.Error())
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Something is wrong with invocation response: " + err.Error())
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}
