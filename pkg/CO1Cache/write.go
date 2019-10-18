package CO1Cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Write(file string, content []byte) {
	err := ioutil.WriteFile(Filename(file), content, 644)

	if err != nil {
		fmt.Printf("Unable to write cache file: %v\n", err)
	}
}

func WriteJSON(file string, content interface{}) {
	b, err := json.Marshal(content)

	if err != nil {
		fmt.Printf("Unable to parse JSON: %v\n", err)
		return
	}

	Write(file, b)
}
