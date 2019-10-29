package repocheck

import (
	"log"
	"net/http"
)

func ErrorHandler(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
	}
}
