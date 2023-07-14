package utils

import (
	"log"
	"os"
	"strconv"
)

func LookupEnvString(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)

	if !exist {
		return defaultValue
	} else {
		return value
	}
}

func LookupEnvInt(key string, defaultValue int) int {
	value, exist := os.LookupEnv(key)
	if !exist {
		return 0
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Key %s must be integer", key)
		return 0
	}

	return valueInt
}
