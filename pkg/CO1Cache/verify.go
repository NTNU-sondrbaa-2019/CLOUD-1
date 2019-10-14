package CO1Cache

import (
	"fmt"
	"os"
	"time"
)

func Verify(file string) bool {
	f, err := os.Stat(Filename(file))

	if err != nil {
		fmt.Printf("Unable to read cache file: %v\n", err)
		return false
	}

	return time.Now().Sub(f.ModTime()) < time.Second * CACHE_DURATION
}
