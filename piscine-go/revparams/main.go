package main

import (
	"os"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		os.Stdout.Write([]byte(os.Args[i] + "\n"))
	}
}
