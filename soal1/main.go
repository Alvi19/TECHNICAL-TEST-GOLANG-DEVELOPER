package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "data.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	var sum, count int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error mengonversi angka:", err)
			return
		}

		sum += num
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file:", err)
		return
	}

	fmt.Println("Total angka pada file:", count)
	fmt.Println("Jumlah semua angka:", sum)
}
