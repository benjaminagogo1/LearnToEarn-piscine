package main

import (
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		os.Stdout.Write([]byte(os.Args[i] + "\n"))
	}
}
