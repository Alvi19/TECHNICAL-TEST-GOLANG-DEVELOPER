package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffleRecursive(n int) {
	if n == 0 {
		return
	}

	randNum := rand.Intn(10) + 1
	fmt.Print(randNum, " ")

	shuffleRecursive(n - 1)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	shuffleRecursive(10)
	fmt.Println()
}
