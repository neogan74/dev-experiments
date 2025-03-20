package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "htto, %s!\n", ps.ByName("name"))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	port := ":8080"
	logger.Info("Starting server", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
