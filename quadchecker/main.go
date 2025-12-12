package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	input := readInput()
	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	height := len(lines)
	width := len(lines[0])

	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matches []string

	for _, q := range quads {
		out, err := exec.Command("./"+q, fmt.Sprintf("%d", width), fmt.Sprintf("%d", height)).Output()
		if err != nil {
			continue
		}

		if bytes.Equal(bytes.TrimSpace(out), bytes.TrimSpace([]byte(input))) {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", q, width, height))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		sort.Strings(matches)
		fmt.Println(strings.Join(matches, " || "))
	}
}

func readInput() string {
	var buf bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf.WriteString(scanner.Text() + "\n")
	}
	return buf.String()
}
