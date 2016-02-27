package main

import (
	"os"
	"fmt"
	"log"
)

const usage string = "usage: spredendb <command>"


func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		err := createDatabase(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
	case "tree":
		fmt.Println("tree command")
	case "status":
		fmt.Println("status command")
	}

	os.Exit(0)
}
