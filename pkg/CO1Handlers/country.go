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
returns a list of species based on a country code and limit parameters
country/{:country_identifier}{?limit={:limit}}
*/
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	//Separates the url to get the country code
	parts := strings.Split(r.URL.Path, "/")
	countryCode := parts[len(parts)-1]

	//Check if this result is in cache
	if CO1Cache.Verify("country" + countryCode) {
		var countryOut CO1Struct.Country
		err := json.Unmarshal(CO1Cache.Read("country" + countryCode), &countryOut)
		if err != nil {
			fmt.Printf("Unable to read country information from cache: %v\n", err)
		}

		//Adds JSON to the header and encodes the output object to JSON
		http.Header.Add(w.Header(), "content-type", "application/json")
		err = json.NewEncoder(w).Encode(countryOut)
		if err != nil {
			http.Error(w, "Could not encode", 400)
		}
	} else {

		//GETs information about the country from RCEU
		resp, err := http.Get("https://restcountries.eu/rest/v2/alpha/" + countryCode)
		if err != nil {
			http.Error(w, "Could not get country, Bad request", 400)
			return
		}

		dec := json.NewDecoder(resp.Body)
		var countryInfo CO1Struct.CountryRCEU

		err = dec.Decode(&countryInfo)
		if err != nil {
			http.Error(w, "Could not decode json", 400)
			return
		}

		//GETs the species that have occurred recently in the country from GBIF
		resp, err = http.Get("http://api.gbif.org/v1/occurrence/search?country=" + countryCode)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}

		dec = json.NewDecoder(resp.Body)
		var speciesInfo CO1Struct.ResultCountryGBIF

		err = dec.Decode(&speciesInfo)
		if err != nil {
			http.Error(w, "Could not decode json", 400)
			return
		}

		//Combines the two results from the get requests into one output object
		var countryOut CO1Struct.Country
		countryOut = CountryFmt(speciesInfo, countryInfo)

		//Write to cache
		CO1Cache.WriteJSON("country" + countryCode, countryOut)

		//Adds JSON to the header and encodes the output object to JSON
		http.Header.Add(w.Header(), "content-type", "application/json")
		err = json.NewEncoder(w).Encode(countryOut)
		if err != nil {
			http.Error(w, "Could not encode", 400)
		}
	}
}

//Combines the variables from the two get requests
func CountryFmt (result CO1Struct.ResultCountryGBIF, country CO1Struct.CountryRCEU) CO1Struct.Country {
	var out CO1Struct.Country
	out.Code = country.Code
	out.CountryName = country.CountryName
	out.CountryFlag = country.CountryFlag

	//Checks if the species and species key are duplicate before appending them
	speciesDuplicate := false
	speciesKeyDuplicate := false
	for i:=0; i<len(result.ResultArray); i++ {
		speciesDuplicate = false
		speciesKeyDuplicate = false
		if i>0 {
			for j:=0; j<len(out.Species); j++ {
				if out.Species[j]==result.ResultArray[i].Species {speciesDuplicate = true}
			}
			for j:=0; j<len(out.SpeciesKey); j++ {
				if out.SpeciesKey[j]==result.ResultArray[i].SpeciesKey {speciesKeyDuplicate = true}
			}
		}
		if !speciesDuplicate {out.Species = append(out.Species, result.ResultArray[i].Species)}
		if !speciesKeyDuplicate {out.SpeciesKey = append(out.SpeciesKey, result.ResultArray[i].SpeciesKey)}
	}

	return out
}