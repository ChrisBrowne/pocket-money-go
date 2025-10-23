package utils_test

import (
	"pocketmoney/internal/utils"
	"testing"
)

func envvar(key string) string {
	switch key {
	case "PORT":
		return "3000"
	default:
		return ""
	}
}

func TestGetEnvIntReturnsCorrectly(t *testing.T) {
	actual := utils.GetEnvInt(envvar, "PORT", 1)
	if actual != 3000 {
		t.Errorf("port should be 3000, got %d", actual)
	}
}

func TestGetEnvIntReturnsDefault(t *testing.T) {
	actual := utils.GetEnvInt(envvar, "INVALID_ENV_VAR", 1)
	if actual != 1 {
		t.Errorf("port should be defaulted to 1, got %d", actual)
	}
}
