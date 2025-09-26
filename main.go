package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func getEnvInt(key string, defaultValue int) int {
	valueStr, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return valueInt
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pocket money, golang edition")
}

func main() {
	port := getEnvInt("PORT", 8080)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("server started successfully", "port", port)

	http.HandleFunc("/", roothandler)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
