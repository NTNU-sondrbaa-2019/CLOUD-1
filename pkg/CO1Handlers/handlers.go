package CO1Handlers

import (
	"net/http"
	"fmt"
)

//This is the default handler for invalid requests
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default handler: invalid request received.")
	http.Error(w, "Invalid Request", http.StatusBadRequest)
}

/*
returns a list of species based on a country code and limit parameters
country/{:country_identifier}{?limit={:limit}}
 */
func HandlerCountry(w http.ResponseWriter, r *http.Request) {

}

/*
provides information on specific species
species/{:speciesKey}
 */
func HandlerSpecies(w http.ResponseWriter, r *http.Request) {

}

/*
Indicates the availability of the services this service depends on
 */
func HandlerDiag(w http.ResponseWriter, r *http.Request) {

}