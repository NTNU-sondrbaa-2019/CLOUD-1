package main

import (
	"CLOUD-O1/pkg/CO1Cache"
	"CLOUD-O1/pkg/CO1Handlers"
	"fmt"
	"net/http"
	"os"
	"log"
)

func main () {

	CO1Cache.Initialize()

	/*

	CACHE TESTING

	type test struct {
		name string
		number int
	}

	heh := &test{
		name: "test",
		number: 123,
	}

	CO1Cache.WriteJSON("test", heh)
	if CO1Cache.Verify("test") {
		fmt.Println(string(CO1Cache.Read("test")))
	}
	*/
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
