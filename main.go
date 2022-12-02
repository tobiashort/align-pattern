package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Line struct {
	Text  string
	Index []int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: align-pattern PATTERN")
		os.Exit(1)
	}

	pattern := regexp.MustCompile(os.Args[1])
	reader := bufio.NewReader(os.Stdin)
	lines := make([]Line, 0)

	for {
		text, err := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if err != nil && err != io.EOF {
			panic(err)
		}

		index := pattern.FindStringIndex(text)
		lines = append(lines, Line{Text: text, Index: index})

		if err == io.EOF {
			break
		}
	}

	maxIndex := 0

	for _, line := range lines {
		if line.Index != nil && line.Index[0] > maxIndex {
			maxIndex = line.Index[0]
		}
	}

	for _, line := range lines {
		if line.Index == nil {
			fmt.Println(line.Text)
		} else {
			index := line.Index[0]
			head := line.Text[:index]
			tail := line.Text[index:]
			padding := strings.Repeat(" ", maxIndex-index)
			fmt.Println(head + padding + tail)
		}
	}
}
