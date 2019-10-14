package CO1Cache

import (
	"encoding/json"
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

func ReadJSON(file string) []map[string]interface{} {
	f, err := ioutil.ReadFile(Filename(file))

	if err != nil {
		fmt.Printf("Unable to read cache file: %v\n", err)
		return nil
	}

	data := []map[string]interface{}{}
	err = json.Unmarshal(f, &data)

	if err != nil {
		fmt.Printf("Unable to read JSON: %v\n", err)
		return nil
	}

	return data
}
