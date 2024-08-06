package main


import (
	"strings"
	"unicode"
)

func wordFrequency(input string) map[string]int {
	counter := make(map[string]int)
	var wordBuilder strings.Builder

	for _, char := range input {
		if unicode.IsPunct(char) || unicode.IsSpace(char) {
			if wordBuilder.Len() > 0 {
				word := strings.ToLower(wordBuilder.String())
				counter[word]++
				wordBuilder.Reset()
			}
		} else {
			wordBuilder.WriteRune(char)
		}
	}

	if wordBuilder.Len() > 0 {
		word := strings.ToLower(wordBuilder.String())
		counter[word]++
	}

	return counter
}


