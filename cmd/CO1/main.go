package main

import (
	"CLOUD-O1/pkg/CO1Cache"
	"fmt"
)

func main () {

	CO1Cache.Initialize()
	fmt.Println("TEST")

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
}
