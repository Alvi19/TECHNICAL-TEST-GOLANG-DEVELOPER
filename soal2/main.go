package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var cleaned string

	for _, r := range s {
		if unicode.IsLetter(r) {
			cleaned += string(unicode.ToLower(r))
		}
	}

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Masukkan kata atau kalimat:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if isPalindrome(input) {
		fmt.Printf("Kata/kalimat yang diinput: \"%s\"\nMerupakan palindrom\n", input)
	} else {
		fmt.Printf("Kata/kalimat yang diinput: \"%s\"\nBukan palindrom\n", input)
	}
}
