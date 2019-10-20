package CO1Handlers

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Struct"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
)

/*
provides information on specific species
species/{:speciesKey}
*/
func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	//Separates the URL to get the species key
	parts := strings.Split(r.URL.Path, "/")
	speciesKey := parts[len(parts)-1]

	//Checks that

	//Check if this result is in cache
	if CO1Cache.Verify("specie" + speciesKey) {
		var speciesOut CO1Struct.Specie
		err := json.Unmarshal(CO1Cache.Read("specie" + speciesKey), &speciesOut)
		if err != nil {
			fmt.Printf("Unable to read species from cache: %v\n", err)
		}

		//Adds JSON to the header and encodes the output object to JSON
		http.Header.Add(w.Header(), "content-type", "application/json")
		err = json.NewEncoder(w).Encode(speciesOut)
		if err!= nil {
			http.Error(w, "Could not encode json", 400)
			return
		}
	} else {

		//GETs the species genetic information
		resp, err := http.Get("http://api.gbif.org/v1/occurrence/search?speciesKey=" + speciesKey)
		if err != nil {
			http.Error(w, "Bad request species", 400)
			return
		}

		if resp.StatusCode != 200 {
			http.Error(w, "Bad response from GBIF, no species with that key", 400)
			return
		}

		dec1 := json.NewDecoder(resp.Body)
		var responsesGBIF CO1Struct.ResultSpeciesGBIF

		err = dec1.Decode(&responsesGBIF)
		if err != nil {
			http.Error(w, "Could not decode species json", 400)
			return
		}

		//GETs the species name and year information
		resp, err = http.Get("http://api.gbif.org/v1/species/" + speciesKey + "/name")
		if err != nil {
			http.Error(w, "Bad Request name", 400)
			return
		}

		dec := json.NewDecoder(resp.Body)
		var namesGBIF CO1Struct.NameGBIF

		err = dec.Decode(&namesGBIF)
		if err != nil {
			http.Error(w, "Could not decode names json", 400)
			return
		}

		//Combines the results from the two GET requests into one output object
		var speciesOut CO1Struct.Specie
		speciesOut = SpeciesFmt(responsesGBIF, namesGBIF)

		//Write to cache
		CO1Cache.WriteJSON("specie" + speciesKey, speciesOut)

		//Adds JSON to the header and encodes the output object to JSON
		http.Header.Add(w.Header(), "content-type", "application/json")
		err = json.NewEncoder(w).Encode(speciesOut)
		if err != nil {
			http.Error(w, "Could not encode json", 400)
			return
		}
	}
}

//Combines the variables from the two get requests
func SpeciesFmt (result CO1Struct.ResultSpeciesGBIF, name CO1Struct.NameGBIF) CO1Struct.Specie {
	var out CO1Struct.Specie
	out.Key = result.ResultArray[0].Key
	out.Kingdom = result.ResultArray[0].Kingdom
	out.Phylum = result.ResultArray[0].Phylum
	out.Order = result.ResultArray[0].Order
	out.Family = result.ResultArray[0].Family
	out.Genus = result.ResultArray[0].Genus
	out.ScientificName = name.ScientificName
	out.CanonicalName = name.CanonicalName
	out.Year = name.Year

	return out
}