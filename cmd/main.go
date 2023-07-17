package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leonardogregoriocs/internal/domain/client"
	"github.com/leonardogregoriocs/internal/domain/transactions"
	"github.com/leonardogregoriocs/internal/endpoints"
	"github.com/leonardogregoriocs/internal/infra/database"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()

	ClientService := client.Service{
		Repository: &database.ClientRepository{Db: db},
	}

	TransactionsService := transactions.Service{
		Repository:       &database.TransactionsRepository{Db: db},
		ClientRepository: &database.ClientRepository{Db: db},
	}

	handler := endpoints.Handler{
		ClientService:       ClientService,
		TransactionsService: TransactionsService,
	}

	r.Post("/accounts", handler.ClientPost)
	r.Get("/accounts/{id}", handler.ClientGetById)
	r.Post("/transactions", handler.TransactionsPost)

	http.ListenAndServe(":8000", r)
}
