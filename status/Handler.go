package status

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// StartTime - Set by init() in main.go
var StartTime time.Time

// Uptime - Returns times in seconds since service start
func Uptime() time.Duration {
	return time.Since(StartTime) / 1000000000
}

type status struct {
	Gitlab   int
	Database int
	Uptime   int
	Version  string
}

// Handler - Comment
func Handler(w http.ResponseWriter, r *http.Request) {
	// Accept only GET requests
	switch r.Method {
	case http.MethodGet:
		var stat status
		resp, err := http.Get("https://git.gvk.idi.ntnu.no/")
		if err != nil {
			log.Fatalln(err)
		}
		stat.Gitlab = resp.StatusCode

		resp, err = http.Get("https://firebase.google.com/")
		if err != nil {
			log.Fatalln(err)
		}
		stat.Database = resp.StatusCode
		stat.Version = "v1"
		stat.Uptime = int(Uptime())

		// Encode new structure to JSON format
		enc, err := json.Marshal(stat)
		if err != nil {
			log.Fatalln(err)
		}

		// Gives JSON response for requests
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(enc)
	default:
		// Methods not allowed - Returns 405
		fmt.Println("HandlerStatus.go: Method not Allowed or Implemented" + r.Method)
		http.Error(w, "Method not allowed: "+r.Method, http.StatusMethodNotAllowed /* 405 */)
	}
}
