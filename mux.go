package main

import (
	"context"
	"go-port/clock"
	"go-port/config"
	"go-port/handler"
	"go-port/store"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	// http.serveMux()は色々な問題点があるので、利用しない
	// mux := http.NewServeMux()
	// 代わりに chi を利用して、routerを記載する！
	mux := chi.NewRouter()
	// http.serveMux の記法は chi でも利用できる
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{DB: db, Repo: &r, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{DB: db, Repo: &r}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux, cleanup, nil
}
