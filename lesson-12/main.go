package main

import (
	_ "net/http/pprof"
	"net/http"
	"os"
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	http.HandleFunc("/stuck", StuckGoroutine)
	http.HandleFunc("/new-file", CreateFile)
	http.ListenAndServe(":6060", nil)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("./test.txt", os.O_CREATE | os.O_CREATE, 0777)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	file.Write([]byte("Some data"))
	w.Write([]byte("Done."))
}

func StuckGoroutine(w http.ResponseWriter, r *http.Request) {
	go stuckGoroutine()
	w.Write([]byte("Done."))
}

func stuckGoroutine() {
	ch := make(chan bool)
	fmt.Println("stuck goroutine")
	<-ch
}
