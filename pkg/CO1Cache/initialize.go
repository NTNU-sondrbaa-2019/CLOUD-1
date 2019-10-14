package CO1Cache

import (
	"fmt"
	"os"
)

func Initialize() {
	err := os.MkdirAll(CACHE_PATH, 755)

	if err != nil {
		fmt.Printf("Unable to create cache directory: %v\n", err)
	}
}
