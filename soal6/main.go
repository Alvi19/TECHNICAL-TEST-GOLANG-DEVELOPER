package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1. FizzBuzz
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// 2. Palindrom Checker
func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// 3. Faktorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 4. Nilai Terbesar dan Terkecil dalam Array
func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	min, max := arr[0], arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

// 5. Menghitung jumlah huruf dan angka dalam string
func countLettersAndDigits(s string) (int, int) {
	letters, digits := 0, 0
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			letters++
		} else if unicode.IsDigit(ch) {
			digits++
		}
	}
	return letters, digits
}

func main() {
	// 1. Menjalankan FizzBuzz
	fmt.Println("=== FizzBuzz ===")
	fizzBuzz()

	// 2. Cek Palindrom
	fmt.Println("\n=== Palindrom Checker ===")
	testStr := "Kasur ini rusak"
	fmt.Printf("Apakah \"%s\" palindrom? %v\n", testStr, isPalindrome(testStr))

	// 3. Faktorial
	fmt.Println("\n=== Faktorial ===")
	n := 5
	fmt.Printf("Faktorial dari %d adalah %d\n", n, factorial(n))

	// 4. Nilai Min dan Max dalam Array
	fmt.Println("\n=== Nilai Min dan Max ===")
	arr := []int{7, 2, 9, 4, 1, 8}
	min, max := findMinMax(arr)
	fmt.Printf("Nilai terkecil: %d, Nilai terbesar: %d\n", min, max)

	// 5. Menghitung Huruf dan Angka dalam String
	fmt.Println("\n=== Hitung Huruf dan Angka ===")
	str := "Golang123"
	letters, digits := countLettersAndDigits(str)
	fmt.Printf("Jumlah huruf: %d, Jumlah angka: %d\n", letters, digits)
}
