package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"time"

	"errors"

	"github.com/urfave/cli"
)

var (
	isDaemon = false
	VERSION  = "1.0.0"
)

var (
	ErrNotRunning    = errors.New("Process is not running")
	ErrUnableToParse = errors.New("Unable to read and parse process id")
	ErrUnableToKill  = errors.New("Unable to kill process")
)

func main() {
	app := cli.NewApp()
	app.Name = "Match Maker"
	app.Usage = "match maker service"
	app.UsageText = "matchmaker [global options]"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:      "run",
			Usage:     "run command",
			UsageText: "main run",
			Action:    run,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "daemon, d",
					Usage:       "daemon flag",
					Destination: &isDaemon,
				},
			},
		},
		{
			Name:      "stop",
			Usage:     "stop command",
			UsageText: "main stop",
			Action:    stop,
		},
	}

	fmt.Println(app.Run(os.Args))
}

func run(ctx *cli.Context) error {
	if isDaemon {
		return runDaemon(ctx)
	}

	// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <-ch
		signal.Stop(ch)
		fmt.Println("Exit command received. Exiting...")

		// this is a good place to flush everything to disk
		// before terminating.
		fmt.Println("Received signal type : ", signalType)

		// remove PID file
		os.Remove(getPidFilePath())

		os.Exit(0)
	}()

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello World!")
	}
}

func stop(ctx *cli.Context) error {
	if _, err := os.Stat(getPidFilePath()); err != nil {
		return ErrNotRunning
	}

	data, err := ioutil.ReadFile(getPidFilePath())
	if err != nil {
		return ErrNotRunning
	}
	ProcessID, err := strconv.Atoi(string(data))

	if err != nil {
		return ErrUnableToParse
	}

	process, err := os.FindProcess(ProcessID)
	if err != nil {
		return ErrUnableToParse
	}
	// remove PID file
	os.Remove(getPidFilePath())

	fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
	// kill process and exit immediately
	err = process.Kill()

	if err != nil {
		return ErrUnableToKill
	}

	fmt.Printf("Killed process ID [%v]\n", ProcessID)
	return nil
}

func runDaemon(ctx *cli.Context) error {
	// check if daemon already running.
	if _, err := os.Stat(getPidFilePath()); err == nil {
		fmt.Println("Already running or /tmp/daemonize.pid file exist.")
		os.Exit(1)
		return nil
	}

	cmd := exec.Command(os.Args[0], "run")
	cmd.Start()
	fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
	savePID(cmd.Process.Pid)
	os.Exit(0)

	return nil
}

func savePID(pid int) {
	file, err := os.Create(getPidFilePath())
	if err != nil {
		log.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		log.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() // flush to disk
}

func getPidFilePath() string {
	return os.Getenv("HOME") + "/daemon.pid"
}
