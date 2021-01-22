package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/pflag"
	"inet.af/tcpproxy"
)

var (
	showHelp    bool
	showVersion bool
	entries     []string
	proxy       tcpproxy.Proxy
)

const (
	version = "0.1.0"
)

type stderrWriter struct{}

func (w stderrWriter) Write(bytes []byte) (int, error) {
	return fmt.Fprint(os.Stderr, string(bytes))
}

func usage() string {
	return "usage of toxy:\n" + pflag.CommandLine.FlagUsages()
}

func main() {
	writer := new(stderrWriter)
	// disable default log format
	log.SetFlags(0)
	log.SetOutput(writer)
	// set writer for pflag
	pflag.CommandLine.SetOutput(writer)
	// command-line flags
	pflag.StringSliceVarP(&entries, "forward", "f", []string{}, "entry to forward on the format <src-ip>:<src-port>-<dst-ip>:<dst-port>")
	pflag.BoolVarP(&showHelp, "help", "h", false, "you're looking at it")
	pflag.BoolVarP(&showVersion, "version", "v", false, "show version")
	pflag.Parse()

	if showVersion {
		log.Print(version)
		return
	}

	if showHelp || len(entries) == 0 {
		log.Fatal(usage())
	}

	for _, e := range entries {
		parts := strings.Split(e, "-")
		src, dst := parts[0], parts[1]
		proxy.AddRoute(src, tcpproxy.To(dst))
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer close(sigs)

	go func() {
		<-sigs
		// ignore any errors as we're about to shut down
		_ = proxy.Close()
	}()

	err := proxy.Start()
	if err != nil {
		_ = proxy.Close()
		log.Fatalf("error starting: %+v\n", err)
	}
	_ = proxy.Wait()
}
