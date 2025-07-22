package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func counter(input string) map[string]int {
	count := make(map[string]int)
	currentWord := strings.Builder{}

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			currentWord.WriteRune(unicode.ToLower(r))
		} else {
			if currentWord.Len() > 0 {
				word := currentWord.String()
				count[word]++
				currentWord.Reset()
			}
		}
	}

	if currentWord.Len() > 0 {
		word := currentWord.String()
		count[word]++
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the text you want to count:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	freq := counter(input)

	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
