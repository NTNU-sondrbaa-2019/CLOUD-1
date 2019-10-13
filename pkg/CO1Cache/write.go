package CO1Cache

import (
	"fmt"
	"io/ioutil"
)

func Write(file string, content string) {
	err := ioutil.WriteFile(filename(file), []byte(content), 0755)
	if err != nil {
		fmt.Printf("Unable to write cache file: %v", err)
	}
}
