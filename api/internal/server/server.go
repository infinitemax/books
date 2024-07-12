package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/infinitemax/books/internal/hello"
	"net/http"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}

	return server
}

func (s *Server) Init() error {
	s.Router.Route("/api", func(r chi.Router) {
		r.Get("/hello", hello.HelloHandler)
	})
	return nil
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen %s\n", err)
		}
	}()
	fmt.Println("server started on port 8080")

	<-ctx.Done()
	fmt.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	fmt.Println("Server exited properly")
	return nil
}
