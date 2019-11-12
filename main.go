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
		digits[1], digits[0],
		colons,
		digits[4], digits[3],
		colons,
		digits[6], digits[2],
	}

	drawDigits(digits)
	fmt.Println()
	drawClock(clock)

}
