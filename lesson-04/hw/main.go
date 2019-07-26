package main

import (
	"os"
	"io"
	"bufio"
	"time"
)

const (
	inputFilename = "./input.txt"
	outputFilename = "./output.txt"
)

type LocalWriter struct {
	writer io.Writer
}

func (lcl *LocalWriter) Write(data []byte) (n int, err error) {
	return lcl.writer.Write([]byte(time.Now().String() + " : " + string(data) + "\n"))
}

func main() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buff := make([]byte, 256)
	stdOutCh := make(chan []byte)
	fileOutCh := make(chan []byte)

	defer close(stdOutCh)
	defer close(fileOutCh)

	outputFile, err := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	fileWriter := bufio.NewWriter(&LocalWriter{outputFile})
	stdoutWriter := bufio.NewWriter(&LocalWriter{os.Stdout})

	go writeDataToWriter(fileOutCh, fileWriter)
	go writeDataToWriter(stdOutCh, stdoutWriter)

	go flushWriter(fileWriter)
	go flushWriter(stdoutWriter)

	for err != io.EOF {
		time.Sleep(100 * time.Millisecond)
		_, err = reader.Read(buff)

		go sendDataToChannel(stdOutCh, buff)
		go sendDataToChannel(fileOutCh, buff)
	}
}

func sendDataToChannel(ch chan <- []byte, d []byte) {
	ch <- d
}

func writeDataToWriter(byteCh <- chan []byte, writer io.Writer) {
	for d := range byteCh {
		writer.Write(d)
	}
}

func flushWriter(writer *bufio.Writer) {
	var err error
	for err != nil {
		time.Sleep(1 * time.Second)
		err = writer.Flush()
	}
}
