package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: analyze <file>")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines, words int64
	stat, _ := file.Stat()
	size := stat.Size()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		for _, w := range splitWords(scanner.Text()) {
			if w != "" {
				words++
			}
		}
	}

	fmt.Printf("File: %s\nSize: %d bytes\nLines: %d\nWords: %d\n", filePath, size, lines, words)
}

func splitWords(s string) []string {
	var res []string
	word := ""
	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' {
			if word != "" {
				res = append(res, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
