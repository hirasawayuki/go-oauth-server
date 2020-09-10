package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hirasawayuki/go-oauth-server/generator"
)

func main() {
	appHandler, cleanup, err := generator.InitializeAppHandler()
	defer cleanup()
	if err != nil {
		log.Fatal("oops, something was too hard", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/authorize", appHandler.AuthorizeHandler.Authorize)

	server := &http.Server{
		Addr:         "127.0.0.1:9001",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	server.ListenAndServe()
}
