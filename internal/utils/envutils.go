package utils

import (
	"strconv"
)

func GetEnvInt(getenvvar func(string) string, key string, defaultValue int) int {
	valueStr := getenvvar(key)

	if valueStr == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return valueInt
}
