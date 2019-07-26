package main

import (
	"fmt"
	"bufio"
	"time"
)

type ConsoleWriter struct {}

func (wr ConsoleWriter) Write(data []byte) (n int, err error) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(string(data))
	return len(data), nil
}

func main() {
	//cnsl :=
	//cnsl.Write([]byte("2"))
	//cnsl.Write([]byte("3"))

	bw := bufio.NewWriter(&ConsoleWriter{})
	//bw := &ConsoleWriter{}
	bw.Write([]byte("1"))
	bw.Write([]byte("2"))
	bw.Write([]byte("3"))
	bw.Write([]byte("4"))
	bw.Write([]byte("5"))
	bw.Write([]byte("6"))
	bw.Write([]byte("7"))
	bw.Flush()
}
