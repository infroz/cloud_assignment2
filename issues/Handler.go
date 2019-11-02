package issues

import (
	"fmt"
	"net/http"
)

// Handler - Comment
func Handler(w http.ResponseWriter, r *http.Request) {
	// Accept only GET requests
	switch r.Method {

	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerCommits.go: Method not Allowed or Implemented" + r.Method)
		http.Error(w, "Method not allowed: "+r.Method, http.StatusMethodNotAllowed /* 405 */)
	}
}
