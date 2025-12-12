package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, arg := range args {
		os.Stdout.Write([]byte(arg + "\n"))
	}
}
