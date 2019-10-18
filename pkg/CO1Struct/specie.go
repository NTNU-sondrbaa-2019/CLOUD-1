package CO1Struct

type Specie struct {
	Key				int			`json:"key"`
	Kingdom			string		`json:"kingdom"`
	Phylum			string		`json:"phylum"`
	Order			string		`json:"order"`
	Family			string		`json:"family"`
	Genus			string		`json:"genus"`
	ScientificName	string		`json:"scientificName"`
	CanonicalName	string		`json:"canonicalName"`
	Year			int			`json:"year"`
}
