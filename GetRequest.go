package repocheck

import (
	"log"
	"net/http"
)

// GetRequest - Comment
func GetRequest(c *http.Client, s string) *http.Response {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
