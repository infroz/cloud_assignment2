package languages

import (
	"encoding/json"
	"log"
	"net/http"
	"repocheck"
	"strconv"
)

func getLanguages(repoID int, auth string) map[string]float32 {
	client := http.DefaultClient
	request := repocheck.API + "projects/" + strconv.Itoa(repoID) + "/languages?private_token=" + auth
	response := repocheck.GetRequest(client, request)

	var tempLang map[string]float32
	err := json.NewDecoder(response.Body).Decode(&tempLang)
	if err != nil {
		log.Fatalln(err)
	}

	return tempLang
}
