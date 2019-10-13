package CO1Cache

import (
	"fmt"
	"io/ioutil"
)

func read(file string) []byte {
	b, err := ioutil.ReadFile(filename(file))
	if err != nil {
		fmt.Printf("Unable to read cache file: %v", err)
	}
	return b
}
