package CO1Cache

import (
	"fmt"
)

func filename(file string) string {
	return fmt.Sprintf("%s%s.cache", CACHE_PATH, file)
}
