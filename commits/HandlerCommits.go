package commits

import (
	"fmt"
	"net/http"
)

/*		* The endpoint should accept GET requests with empty payload
 *		* If not specified, the parameter limit should return 5 repositories
 *		*	If not specified in the auth parameter, the request should occur
 *			without authentication
 *
 *		url: /repocheck/v1/commits?limit=[]&auth=[authentication token]
 */

// HandlerCommits - Handler for any requests to /v1/commits
func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	// Accept only GET requests
	switch r.Method {
	case http.MethodGet:
		// Checked for GET method...

	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerCommits.go: Method not Allowed")
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed /* 405 */)
	}
}
