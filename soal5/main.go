package main

import (
	"fmt"
	"regexp"
	"time"
)

func convertTo24HourFormat(time12 string) (string, error) {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2]):([0-5][0-9]):([0-5][0-9])\s?(AM|PM)$`)
	if !re.MatchString(time12) {
		return "", fmt.Errorf("Invalid input format")
	}
	t, err := time.Parse("03:04:05 PM", time12)
	if err != nil {
		return "", fmt.Errorf("Invalid input format")
	}

	return t.Format("15:04:05"), nil
}

func main() {
	var timePart, period string
	fmt.Print("Masukkan waktu dalam format 12 jam (HH:MM:SS AM/PM): ")
	fmt.Scanf("%s %s", &timePart, &period)

	input := timePart + " " + period
	convertedTime, err := convertTo24HourFormat(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Waktu dalam format 24 jam:", convertedTime)
}
