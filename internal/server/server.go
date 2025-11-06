package server

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var log *slog.Logger

func init() {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func handleSomething(commandChan chan PocketMoneyCommand) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			respCh := make(chan []Child)
			cmd := GetKidsPocketMoneyCommand{Resp: respCh}
			commandChan <- cmd
			res := <-respCh
			fmt.Printf("Kids: %s", res)
			fmt.Fprintf(w, "Pocket money, golang edition, mooo %s", res)
		},
	)
}

func roothandler(commandChan chan PocketMoneyCommand) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		respCh := make(chan []Child)

		cmd := GetKidsPocketMoneyCommand{Resp: respCh}
		commandChan <- cmd
		res := <-respCh

		names := make([]string, len(res))
		for i, child := range res {
			names[i] = child.Name
		}
		all := strings.Join(names, ", ")

		fmt.Fprintf(w, "Pocket money, golang edition, mooo %s", all)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func addRoutes(mux *http.ServeMux, commandChan chan PocketMoneyCommand) {
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/", roothandler(commandChan))
	mux.HandleFunc("/some", handleSomething(commandChan).ServeHTTP)
}

func AppHandler(commandChan chan PocketMoneyCommand) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, commandChan)

	var handler http.Handler = mux
	return handler
}

func Run(config *Config) error {
	log.Info("starting server", "port", config.Port)

	var pocketMoneyManagerCommandChannel = make(chan PocketMoneyCommand)
	go pocketMoneyManager(pocketMoneyManagerCommandChannel)

	handler := AppHandler(pocketMoneyManagerCommandChannel)
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
