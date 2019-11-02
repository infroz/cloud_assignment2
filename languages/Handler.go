package languages

import (
	"fmt"
	"net/http"
)

// Handler - Accepts requests for /v1/languages/
func Handler(w http.ResponseWriter, r *http.Request) {
	// Accept only GET requests
	switch r.Method {
	case http.MethodGet:
		// Below Code is executed if Method is GET

	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerCommits.go: Method not Allowed")
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed /* 405 */)
	}
}
