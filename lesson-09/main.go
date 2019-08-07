package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fmt"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	select {
	case <-req.Context().Done():
		fmt.Println("connection gone")
		w.Write([]byte("Conn gone"))
	case <-time.After(3 * time.Second):
		w.Write([]byte("Some data"))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", Handler).Methods("GET")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Println("server started")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("received stop signal")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("call to shutdown")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Println("Server stopped")
	}
	time.Sleep(5 * time.Second)
}
