package server

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHttpServer(host, port string) *http.Server {
	app := newHttpServer()
	router := chi.NewRouter()
	router.Post("/produce", app.handleProduce())
	router.Get("/consume", app.handleConsume())
	return &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: router,
	}
}
