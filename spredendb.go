package main

import "os"
import "fmt"

const usage string = "usage: spredendb <command>"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		fmt.Println("create command")
	case "tree":
		fmt.Println("tree command")
	case "status":
		fmt.Println("status command")
	}

	os.Exit(0)
}
