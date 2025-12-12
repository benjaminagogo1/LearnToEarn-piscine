package main

import (
	"os"
)

func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, false
		}
		n = n*10 + int(ch-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]

	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	var output []rune
	for _, arg := range args {
		n, ok := atoi(arg)
		if !ok || n < 1 || n > 26 {
			output = append(output, ' ')
			continue
		}
		if upper {
			output = append(output, rune('A'+n-1))
		} else {
			output = append(output, rune('a'+n-1))
		}
	}

	if len(output) > 0 {
		os.Stdout.Write([]byte(string(output)))
		os.Stdout.Write([]byte("\n"))
	}
}
