package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"ad-api/config"
)

type Server struct {
	Srv    *http.Server
	notify chan error
	db     *sql.DB
}

func New(cnf *config.Config, router *http.ServeMux, db *sql.DB) *Server {
	server := &Server{
		Srv: &http.Server{
			Addr:    ":" + cnf.App_Port,
			Handler: router,
		},
		notify: make(chan error, 1),
		db:     db,
	}
	server.start()

	return server
}

func (s *Server) start() {
	log.Printf("server on http://localhost%v/\n", s.Srv.Addr)
	go func() {
		s.notify <- s.Srv.ListenAndServe()
		close(s.notify)
		log.Printf("notify chan")
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	s.db.Close()
	return s.Srv.Shutdown(ctx)
}
