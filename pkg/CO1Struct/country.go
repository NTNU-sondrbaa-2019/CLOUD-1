package CO1Struct

type ResultCountryGBIF struct {
	ResultArray		[]CountryGBIF	`json:"results"`
}

type CountryGBIF struct {
	Species 		string		`json:"species"`
	SpeciesKey		int			`json:"speciesKey"`
}

type CountryRCEU struct {
	Code 			string		`json:"alpha2Code"`
	CountryName 	string		`json:"name"`
	CountryFlag 	string		`json:"flag"`
}

type Country struct {
	Code 			string		`json:"code"`
	CountryName 	string		`json:"countryname"`
	CountryFlag 	string		`json:"countryflag"`
	Species 		[]string	`json:"species"`
	SpeciesKey 		[]int		`json:"speciesKey"`
}