package CO1Struct

type Diagnostics struct {
	Gbif			int			`json:"gbif"`
	Restcountries	int			`json:"restcountries"`
	Version			string		`json:"version"`
	Uptime			int			`json:"uptime"`
}
