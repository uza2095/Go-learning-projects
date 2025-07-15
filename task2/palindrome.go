package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	
	var input string
	fmt.Println("Enter the string you want to test")
	fmt.Scanln(&input)
	
	input = strings.ToLower(input)
	var cleaned strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned.WriteRune(r)
		}
	}
	cleanStr := cleaned.String()

	
	isPalindrome := true
	for i := 0; i < len(cleanStr)/2; i++ {
		if cleanStr[i] != cleanStr[len(cleanStr)-1-i] {
			isPalindrome = false
			break
		}
	}

	fmt.Printf("Is '%s' a palindrome? %t\n", input, isPalindrome)
}
