package languages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"repocheck/commits"
	"repocheck/webhooks"
	"sort"
	"strconv"
	"time"
)

type lang struct {
	Languages map[string]float32 `json:"languages"`
}

// Handler - Comment
func Handler(w http.ResponseWriter, r *http.Request) {
	// Below is payload handling functions, did not work properly

	/*// Payload handling
	var LangPayload []string // Is used to filter out things not wanted
	payload := false
	//payload := false
	if r.Body != nil {
		var LangPayload []string
		err := json.NewDecoder(r.Body).Decode(&LangPayload)
		if err != nil {
			log.Fatalln(err)
		}
		payload = true
	} else {
		fmt.Fprintf(w, "Body is actually nil")
	}
	*/

	// Accept only GET requests
	switch r.Method {
	case http.MethodGet:
		// Code to be executed when GET request
		Auth := r.URL.Query().Get("auth")
		Limit := r.URL.Query().Get("limit")
		if Limit == "" {
			Limit = "5"
		}

		// Get RepoID's
		repoID := commits.GetProjects(Auth)

		var languages []map[string]float32
		// Checks all repos for languages
		for i := range repoID {
			languages = append(languages, getLanguages(repoID[i].ID, Auth))
		}

		type re struct {
			Lang  string
			Count int
		}

		var response []re

		// Iterates through all repos
		for i := range languages {
			// Iterates through specific repo languages
			for n := range languages[i] {
				/* Below is payload handling functions, did not work properly
				if payload {
					isOk := false
					for p := range LangPayload {
						if n == LangPayload[p] {
							isOk = true
						}
					}
					if !isOk {
						break // Will not add this element
					}
				}*/
				exists := false
				// Checks if language exists in response
				for k := range response {
					if response[k].Lang == n {
						exists = true

						// Increases count if it exists
						response[k].Count++
					}
				}
				if !exists {
					var tmp re
					tmp.Count = 1
					tmp.Lang = n
					response = append(response, tmp)
				}

			}
		}

		// Sorts response
		sort.Slice(response, func(i, j int) bool { return response[j].Count < response[i].Count })

		// Limits response
		var limitedRes []re
		limitInt, err := strconv.Atoi(Limit)
		if err != nil {
			log.Fatalln(err)
		}
		for i := 0; i < limitInt; i++ {
			limitedRes = append(limitedRes, response[i])
		}

		var onlyLang []string
		for i := range limitedRes {
			onlyLang = append(onlyLang, limitedRes[i].Lang)
		}
		// Encode new structure to JSON format
		enc, err := json.Marshal(onlyLang)
		if err != nil {
			log.Fatalln(err)
		}

		// Gives JSON response for requests
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(enc)

		webhooks.URLCaller("languages", "limit="+Limit+" and auth="+Auth, time.Now())
	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerCommits.go: Method not Allowed or Implemented" + r.Method)
		http.Error(w, "Method not allowed: "+r.Method, http.StatusMethodNotAllowed /* 405 */)
	}
}
