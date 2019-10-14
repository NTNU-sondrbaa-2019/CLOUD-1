package CO1Cache

import (
	"fmt"
)

func Filename(file string) string {
	return fmt.Sprintf("%s%s.json", CACHE_PATH, file)
}
