package repocheck

import (
	"fmt"
	"net/http"
)

// HandlerNil - Returns invalid request if nothing is requested
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerNil.go: Invalid request received.")
	http.Error(w, "Invalid request", http.StatusBadRequest)
}
