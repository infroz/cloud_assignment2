package repocheck

import (
	"log"
	"net/http"
)

// ErrorHandler - err is error, s is message, w is http response
func ErrorHandler(err error, w http.ResponseWriter, s string) {
	if s == nil {
		s = "Error" // standard message
	}
	if err != nil {
		log.Fatalln(err)
		http.Error(w, s, http.StatusBadRequest)
	}
}
