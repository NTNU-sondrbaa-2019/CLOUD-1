package CO1Cache

import (
	"fmt"
	"io/ioutil"
)

func Read(file string) []byte {
	b, err := ioutil.ReadFile(Filename(file))

	if err != nil {
		fmt.Printf("Unable to read cache file: %v\n", err)
		return nil
	}

	return b
}
