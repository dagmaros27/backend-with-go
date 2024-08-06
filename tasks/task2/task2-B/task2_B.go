package main

import (
	"strings"
	"unicode"
)

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


func isPalindrome(str string)bool {
	var wordBuilder strings.Builder

	for _, char := range str {
		if unicode.IsPunct(char) || unicode.IsSpace(char) {
				continue
			}
		
			wordBuilder.WriteRune(char)
		}
	
	word := strings.ToLower(wordBuilder.String())
	reverse := reverse(word)
	
	return word == reverse
		
}

