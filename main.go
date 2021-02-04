package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/codingconcepts/see/pkg"
)

func main() {
	seconds := flag.Uint("n", 5, "number of seconds to wait between executions")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("need at least 1 argument to run")
	}
	if *seconds < 1 {
		log.Fatal("n must be at least 1")
	}

	termListen(args, *seconds)

	clear(args, *seconds)
	for {
		outBuf, errBuf := &bytes.Buffer{}, &bytes.Buffer{}

		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = outBuf
		cmd.Stderr = errBuf

		if err := cmd.Run(); err != nil {
			log.Print(err)
		}

		clear(args, *seconds)
		if errBuf.Len() > 0 {
			io.Copy(os.Stderr, errBuf)
		} else {
			io.Copy(os.Stdout, outBuf)
		}

		time.Sleep(time.Second * time.Duration(*seconds))
	}
}

func clear(args []string, seconds uint) {
	pkg.Clear()

	fmt.Printf("Running %s\n", strings.Join(args, " "))
	fmt.Printf("Every: %ds\n\n", seconds)
}

func termListen(args []string, seconds uint) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		clear(args, seconds)
		os.Exit(0)
	}()
}
