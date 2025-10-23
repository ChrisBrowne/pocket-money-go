package server

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

var log *slog.Logger

func init() {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pocket money, golang edition, mooo")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/", roothandler)
}

func AppHandler() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)

	var handler http.Handler = mux
	return handler
}

func Run(config *Config) error {
	log.Info("starting server", "port", config.Port)

	handler := AppHandler()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("localhost", strconv.Itoa(config.Port)),
		Handler: handler,
	}

	// this is a blocking call, unless there is an error
	err := httpServer.ListenAndServe()

	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}

type Config struct {
	Port int
}
