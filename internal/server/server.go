package server

import (
	"embed"
	"fmt"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

//go:embed templates/*.html
var templateFS embed.FS

var t *template.Template

var log *slog.Logger

func init() {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	log.Info("loadingTemplates")
	var err error
	t, err = template.ParseFS(templateFS, "templates/index.html")

	if err != nil {
		log.Error("could not read templates")
		os.Exit(1)
	}
}

func handleSomething(commandChan chan PocketMoneyCommand) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			respCh := make(chan []Child)
			cmd := GetKidsPocketMoneyCommand{Resp: respCh}
			commandChan <- cmd
			res := <-respCh
			fmt.Printf("Kids: %s", strconv.Itoa(len(res)))
			fmt.Fprintf(w, "Pocket money, golang edition, kid count %s", strconv.Itoa(len(res)))
		},
	)
}

func roothandler(commandChan chan PocketMoneyCommand) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		respCh := make(chan []Child)
		cmd := GetKidsPocketMoneyCommand{Resp: respCh}
		commandChan <- cmd
		res := <-respCh

		err := t.ExecuteTemplate(w, "index.html", res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
	store := NewInMemoryChildStore()
	store.SetChild(Child{Name: "Elizabeth", Balance: 5})
	store.SetChild(Child{Name: "Matilda", Balance: 4})
	store.SetChild(Child{Name: "Joseph", Balance: 4})
	go PocketMoneyManager(pocketMoneyManagerCommandChannel, store)

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
