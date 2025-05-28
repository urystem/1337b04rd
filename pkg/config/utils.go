package config

import (
	"log"
	"os"
	"strconv"
)

func mustGetEnvString(key string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Fatalln(key, "not seted")
	}
	return str
}

func mustGetEnvInt(key string) int {
	str := os.Getenv(key)
	if len(str) == 0 {
		log.Fatalln(key, "not seted")
	}

	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(key, "invalid", str)
	}
	return n
}
