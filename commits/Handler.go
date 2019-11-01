package commits

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		var store []project
		store = getProjects() // Gets all project id's

		// Encode new structure to JSON format
		enc, err := json.Marshal(store)
		if err != nil {
			log.Fatalln(err)
		}

		// Gives JSON response for requests
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(enc)

	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerCommits.go: Method not Allowed")
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed /* 405 */)
	}
}
