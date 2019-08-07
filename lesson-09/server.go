package main

import (
	"log"
	"net/http"

	"time"

	"github.com/TV4/graceful"
)

func main() {
	hs := &http.Server{Addr: ":8080", Handler: &server{}}

	go graceful.Shutdown(hs)

	log.Printf("Listening on http://0.0.0.0%s\n", hs.Addr)

	if err := hs.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("Hello!"))
}
