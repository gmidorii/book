package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gmidorii/book/server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	portEnv     = "BOOK_SER_PORT"
	defaultPort = "8088"
)

func serve() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/ping", handler.Ping)
	r.Get("/books", handler.Books)

	p := os.Getenv(portEnv)
	if p == "" {
		p = defaultPort
	}

	log.Printf("PORT: %v\n", p)
	return http.ListenAndServe(fmt.Sprintf(":%v", p), r)
}

func main() {
	if err := serve(); err != nil {
		log.Fatalf("crash server %v", err)
	}
}
