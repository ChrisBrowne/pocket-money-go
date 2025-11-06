package main

import (
	"log/slog"
	"os"
	"pocketmoney/internal/server"
	"pocketmoney/internal/utils"
)

var log *slog.Logger

func init() {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func constructConfig(getenvvar func(string) string) *server.Config {
	return &server.Config{
		Port: utils.GetEnvInt(getenvvar, "PORT", 8080),
	}
}

func main() {
	config := constructConfig(os.Getenv)
	if err := server.Run(config); err != nil {
		log.Error("run returned an error", "error", err)
		os.Exit(1)
	}
}
