package commits

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"repocheck/webhooks"
	"sort"
	"strconv"
	"time"
)

/*		* The endpoint should accept GET requests with empty payload
 *		* If not specified, the parameter limit should return 5 repositories
 *		*	If not specified in the auth parameter, the request should occur
 *			without authentication
 *
 *		url: /repocheck/v1/commits?limit=[]&auth=[authentication token]
 */

// Handler - Handler for any requests to /v1/commits
func Handler(w http.ResponseWriter, r *http.Request) {
	// Accept only GET requests
	switch r.Method {
	case http.MethodGet:
		// Below Code is executed if Method is GET

		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = "5"
		}

		Auth := r.URL.Query().Get("auth")

		var projectStore []ReposTmp

		projectStore = GetProjects(Auth) // Gets all project id's

		var response commits

		for i := range projectStore {
			// Move data to repos
			var tmp repos
			tmp.Repository = projectStore[i].Path
			tmp.Commits = getCommits(projectStore[i].ID, Auth)

			response.Repos = append(response.Repos, tmp)
		}

		// Sorts response
		sort.Slice(response.Repos, func(i, j int) bool { return response.Repos[j].Commits < response.Repos[i].Commits })

		// Set Auth Bool
		if Auth == "" {
			response.Auth = false
		} else {
			response.Auth = true
		}

		limitTmp, err := strconv.Atoi(limit)
		if err != nil {
			log.Fatalln(err)
		}

		var responseLimited commits
		responseLimited.Auth = response.Auth
		for i := 0; i < limitTmp; i++ {
			responseLimited.Repos = append(responseLimited.Repos, response.Repos[i])
		}
		// Encode new structure to JSON format
		enc, err := json.Marshal(responseLimited)
		if err != nil {
			log.Fatalln(err)
		}

		// Gives JSON response for requests
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(enc)

		webhooks.URLCaller("commit", "limit="+limit+" and auth="+Auth, time.Now())
	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerCommits.go: Method not Allowed")
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed /* 405 */)
	}
}
