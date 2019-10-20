package CO1Handlers

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Struct"
	"encoding/json"
	"net/http"
	"time"
)

var Start time.Time
var Version string

/*
Indicates the availability of the services this service depends on
*/
func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	var diagOut CO1Struct.Diagnostics
	resp, err := http.Get("http://api.gbif.org/v1/")
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	diagOut.Gbif = resp.StatusCode

	resp, err = http.Get("http://restcountries.eu/")
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	diagOut.Restcountries = resp.StatusCode

	diagOut.Version = Version

	elapsed := time.Since(Start)
	diagOut.Uptime = elapsed.String()

	http.Header.Add(w.Header(), "content-type", "application/json")
	err = json.NewEncoder(w).Encode(diagOut)
	if err!= nil {
		http.Error(w, "Could not encode json", 400)
		return
	}
}

func StartUptime () {
	Start = time.Now()
}

func SetVersion (v string) {
	Version = v
}