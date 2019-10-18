package CO1Struct

type Country struct {
	Code 			string		`json:"code"`
	CountryName 	string		`json:"countryname"`
	CountryFlag 	string		`json:"countryflag"`
	Species 		[][]byte	`json:"species"`
	SpeciesKey 		[]int		`json:"speciesKey"`
}
