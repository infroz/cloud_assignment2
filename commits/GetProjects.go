package commits

import (
	"encoding/json"
	"log"
	"net/http"
	"repocheck"
	"strconv"
)

func GetProjects(auth string) []ReposTmp {
	var storeTmp []ReposTmp

	request := repocheck.API + "projects?per_page=100&private_token=" + auth

	client := http.DefaultClient
	response := repocheck.GetRequest(client, request)
	// Amount of pages to loop
	pages, err := strconv.Atoi(response.Header.Get("X-Total-Pages"))
	if err != nil {
		log.Fatalln(err)
	}

	// Gets all data from each page
	for i := 1; i <= pages; i++ {
		var tmp []ReposTmp // Temporary stores id to append to store
		response = repocheck.GetRequest(client, request+"&page="+strconv.Itoa(i))
		err = json.NewDecoder(response.Body).Decode(&tmp)
		storeTmp = append(storeTmp, tmp...)
	}
	return storeTmp
}
