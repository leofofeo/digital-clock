package main

import "fmt"

func drawDigits(digits [10]placeholder) {
	for line := 0; line < 5; line++ {
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
}

func drawClock(clock [8]placeholder) {
	for line := range clock[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line], "  ")
		}
		fmt.Println()
	}
}
