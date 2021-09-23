package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			WriteTimeout: 20 * time.Second,
			ReadTimeout:  20 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	log.Printf("[server started on port:%s...]", s.httpServer.Addr)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	log.Printf("[server stoped...]")
	return nil
}
