package main

import "fmt"

type placeholder [5]string

func main() {

	digits := getDigits()

	for line := 0; line < 5; line++ {
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}

}
