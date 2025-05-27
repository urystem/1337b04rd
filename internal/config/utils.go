package config

import (
	"os"
	"strconv"
)

func mustGetEnvString(key string) string {
	str := os.Getenv(key)
	if str == "" {
		return ""
	}
	return str
}

func mustGetEnvInt(key string) int {
	if str := os.Getenv(key); len(str) == 0 {
		return 0
	} else if n, err := strconv.Atoi(str); err != nil {
		return 0
	} else {
		return n
	}
}
