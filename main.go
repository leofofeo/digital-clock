package main

import (
	"fmt"
	"time"
)

type placeholder [5]string

func main() {

	digits := getDigits()

	now := time.Now()

	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	colons := getColons()

	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		colons,
		digits[min/10], digits[min%10],
		colons,
		digits[sec/10], digits[sec%10],
	}

	// drawDigits(digits)
	fmt.Println()
	drawClock(clock)

}
