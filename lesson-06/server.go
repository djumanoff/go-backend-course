package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

type FibNumber struct {
	Current int
	Prev int
	Next int
}

func main() {
	http.HandleFunc("/fibbonaci/4", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
