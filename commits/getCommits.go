package commits

import (
	"net/http"
	"repocheck"
	"strconv"
)

func getCommits(repoID int, auth string) int {
	var commits int

	request := repocheck.API + "projects/" + strconv.Itoa(repoID) + "/repository/commits?private_token=" + auth
	client := http.DefaultClient
	response := repocheck.GetRequest(client, request)

	commits, err := strconv.Atoi(response.Header.Get("X-Total"))
	if err != nil {
		//log.Fatalln("Does not exist " + strconv.Itoa(repoID))
		//log.Fatal(err)
		commits = 0
	}
	return commits
}
