package CO1Cache

import (
	"fmt"
	"os"
	"strconv"
)

func Initialize() {
	err := os.MkdirAll(_CachePath(), 0766)

	if err != nil {
		fmt.Printf("Unable to create cache directory: %v\n", err)
	}
}

func Filename(file string) string {
	return fmt.Sprintf(_CacheFormat(), _CachePath(), file)
}

func _CachePath() string {
	cache_path := os.Getenv("CACHE_PATH")

	if cache_path == "" {
		cache_path = CACHE_PATH
	}

	return cache_path
}

func _CacheDuration() int {
	cache_duration, err := strconv.Atoi(os.Getenv("CACHE_DURATION"))

	if err != nil {
		cache_duration = CACHE_DURATION
	}

	return cache_duration
}

func _CacheFormat() string {
	cache_format := os.Getenv("CACHE_FORMAT")

	if cache_format == "" {
		cache_format = CACHE_FORMAT
	}

	return cache_format
}
