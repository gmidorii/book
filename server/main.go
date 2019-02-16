package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gmidorii/book/server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

const (
	portEnv     = "BOOK_SER_PORT"
)

func serve() error {
	config, err := New()
	if err != nil {
		return errors.Wrap(err, "failed create config")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/ping", handler.Ping)
	r.Get("/books", handler.Books)

	log.Printf("PORT: %v\n", config.Constant.PORT)
	return http.ListenAndServe(fmt.Sprintf(":%v", config.Constant.PORT), r)
}

func main() {
	if err := serve(); err != nil {
		log.Fatalf("crash server %v", err)
	}
}
