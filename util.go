package main

import (
	"os"
	"strings"
)

func wrap(text string) []string {
	var result []string

	paragraphs := strings.Split(text, "\n")

	for _, paragraph := range paragraphs {
		if paragraph == "" {
			result = append(result, "")
			continue
		}

		words := strings.Fields(paragraph)
		if len(words) == 0 {
			continue
		}

		currentLine := words[0]

		for _, word := range words[1:] {
			if len(currentLine)+1+len(word) <= 80 {
				currentLine += " " + word
			} else {
				result = append(result, currentLine)
				currentLine = word
			}
		}

		result = append(result, currentLine)
	}

	return result
}

func getQuery(filename string) string {
	bytes, _ := os.ReadFile(filename)
	return string(bytes)
}
