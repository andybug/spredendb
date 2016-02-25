package main

import (
	"os"
	"fmt"
	"github.com/op/go-logging"
)

// logging setup
var log = logging.MustGetLogger("logger")

const usage string = "usage: spredendb <command>"


func initLogging() {
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc}: %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backendLeveled, backendFormatter)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	initLogging()

	switch os.Args[1] {
	case "create":
		err := createDatabase(os.Args[2], os.Args[3])
		if err != nil {
			log.Critical(err)
		}
	case "tree":
		fmt.Println("tree command")
	case "status":
		fmt.Println("status command")
	}

	os.Exit(0)
}
