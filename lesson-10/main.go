package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	srv             *http.Server
	stopped         bool
	router          *mux.Router
	gracefulTimeout time.Duration
}

func New(srv *http.Server) *Server {
	router := mux.NewRouter()
	srv.Handler = router
	graceTimeout := 5 * time.Second

	return &Server{
		srv,
		false,
		router,
		graceTimeout,
	}
}

func (srv *Server) Shutdown() error {
	srv.stopped = true
	ctx, cancel := context.WithTimeout(context.Background(), srv.gracefulTimeout)
	defer cancel()

	time.Sleep(srv.gracefulTimeout)

	return srv.srv.Shutdown(ctx)
}

func (srv *Server) HandleFunc(method, path string, f func(http.ResponseWriter, *http.Request)) {
	srv.router.HandleFunc(path, srv.makeHandlerFunc(f)).Methods(method)
}

func (srv *Server) ListenAndServe() error {
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop

		if err := srv.Shutdown(); err != nil {
			log.Printf("Error: %v\n", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	return srv.srv.ListenAndServe()
}

func (srv *Server) makeHandlerFunc(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if srv.stopped {
			w.WriteHeader(503)
			return
		}
		select {
		case <-r.Context().Done():
			w.WriteHeader(503)
		default:
			f(w, r)
		}
	}
}

func main() {
	srv := New(&http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	})

	srv.HandleFunc("GET", "/", func(w http.ResponseWriter, request *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte("Some data"))
	})

	srv.ListenAndServe()
}
