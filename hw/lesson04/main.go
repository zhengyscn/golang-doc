package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letterCount() {
	str := "how are you! you are welcome to beijing."
	var wordCount map[string]int = make(map[string]int, 10)
	for i := 0; i < len(str); i++ {
		if unicode.IsLetter(rune(str[i])) {
			if count, ok := wordCount[string(str[i])]; ok {
				wordCount[string(str[i])] = count + 1
			} else {
				wordCount[string(str[i])] = 1
			}
		}
	}
	fmt.Println(wordCount)
}

func wordCount() {
	str := "how are you! you are welcome to beijing."
	var wordCount map[string]int = make(map[string]int, 10)
	for _, word := range strings.Fields(str) {
		if !unicode.IsLetter(rune(word[len(word)-1])) {
			word = word[:len(word)-1]
		}
		if count, ok := wordCount[word]; ok {
			wordCount[word] = count + 1
		} else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)
}

func main() {
	letterCount()
	wordCount()
}
