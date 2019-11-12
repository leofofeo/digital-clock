package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type placeholder [5]string

func main() {
	digits := getDigits()

	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		var colon placeholder
		if sec%2 == 0 {
			colon = placeholder{
				"     ",
				"     ",
				"     ",
				"     ",
				"     ",
			}
		} else {
			colon = getColon()
		}

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		// drawDigits(digits)
		drawClock(clock)
		fmt.Println()

		time.Sleep(time.Second)
	}

}
