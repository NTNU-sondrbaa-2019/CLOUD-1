package CO1Struct

type ResultSpeciesGBIF struct {
	ResultArray		[]SpeciesGBIF	`json:"results"`
}

type SpeciesGBIF struct {
	Key				int			`json:"speciesKey"`
	Kingdom			string		`json:"kingdom"`
	Phylum			string		`json:"phylum"`
	Order			string		`json:"order"`
	Family			string		`json:"family"`
	Genus			string		`json:"genus"`
}

type NameGBIF struct {
	ScientificName	string		`json:"scientificName"`
	CanonicalName	string		`json:"canonicalName"`
	Year			string		`json:"bracketYear"`
}

type Specie struct {
	Key				int			`json:"key"`
	Kingdom			string		`json:"kingdom"`
	Phylum			string		`json:"phylum"`
	Order			string		`json:"order"`
	Family			string		`json:"family"`
	Genus			string		`json:"genus"`
	ScientificName	string		`json:"scientificName"`
	CanonicalName	string		`json:"canonicalName"`
	Year			string		`json:"year"`
}