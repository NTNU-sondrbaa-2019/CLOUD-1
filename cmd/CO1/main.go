package main

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main () {

	CO1Handlers.StartUptime()
	CO1Handlers.SetVersion("v1.0")

	CO1Cache.Initialize()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", CO1Handlers.HandlerNil)
	http.HandleFunc("/conservation/v1/country/", CO1Handlers.HandlerCountry)
	http.HandleFunc("/conservation/v1/species/", CO1Handlers.HandlerSpecies)
	http.HandleFunc("/conservation/v1/diag/", CO1Handlers.HandlerDiag)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}