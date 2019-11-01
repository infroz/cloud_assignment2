package commits

import (
	"encoding/json"
	"log"
	"net/http"
	"repocheck"
	"strconv"
)

func getProjects() []project {
	var store []project
	req := "projects?per_page=100&private_token="
	client := http.DefaultClient
	resp := repocheck.GetRequest(client, repocheck.API+req+repocheck.AuthToken)
	pages, err := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pages)
	for i := 1; i <= pages; i++ {
		resp := repocheck.GetRequest(client, repocheck.API+req+repocheck.AuthToken+"&page="+strconv.Itoa(i))
		var storeTmp []project
		err = json.NewDecoder(resp.Body).Decode(&storeTmp)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Added page: " + strconv.Itoa(i))
		store = append(store, storeTmp...)
	}

	// resp.Header.Get("X-Total")
	return store
}
