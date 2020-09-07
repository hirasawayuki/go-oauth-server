package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hirasawayuki/go-oauth-server/controller"
	"github.com/hirasawayuki/go-oauth-server/handler"
	"github.com/hirasawayuki/go-oauth-server/registry"
	"github.com/hirasawayuki/go-oauth-server/repository"
	"github.com/hirasawayuki/go-oauth-server/usecase"
)

var appHandler handler.AppHandler
var store registry.Store

func init() {
	store, err := registry.RetryConnect(registry.NewMySQLConnection, 10)
	if err != nil {
		log.Fatal(err)
	}
	clientRepository := &repository.ClientRepository{
		Store: store,
	}
	clientService := &service.ClientService{
		ClientRepository: clientRepository,
	}
	clientUsecase := &usecase.ClientUsecase{
		ClientService: clientService,
	}
	clientController := &controller.ClientController{
		ClientUsecase: clientUsecase,
	}
	authorizeHandler := &handler.AuthorizeHandler{
		ClientController: clientController,
	}
	appHandler = handler.AppHandler{
		AuthorizeHandler: authorizeHandler,
	}
}

func main() {
	defer store.DB.Close()
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
