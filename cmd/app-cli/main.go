package main

import (
	"fmt"
	"os"

	"github.com/dyptan-io/go-mono/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		usage(os.Stderr)
		os.Exit(2)
	}

	switch os.Args[1] {
	case "version":
		fmt.Println("app-cli", service.Version)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		usage(os.Stderr)
		os.Exit(2)
	}
}

func usage(w *os.File) {
	fmt.Fprintln(w, "usage: app-cli <command>")
	fmt.Fprintln(w, "commands:")
	fmt.Fprintln(w, "  version    print app-cli version")
}
