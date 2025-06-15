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
	str := mustGetEnvString(key)
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(key, "invalid", str)
	}
	return n
}

func mustGetBoolean(key string) bool {
	switch mustGetEnvString(key) {
	case "true":
		return true
	case "false":
		return false
	default:
		log.Fatalln(key, "invalid")
		return false
	}
}
