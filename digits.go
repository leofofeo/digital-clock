package main

func getDigits() [10]placeholder {
	zero := placeholder{
		"█████",
		"█   █",
		"█   █",
		"█   █",
		"█████",
	}

	one := placeholder{
		" ███ ",
		" ███ ",
		" ███ ",
		" ███ ",
		" ███ ",
	}

	two := placeholder{
		"█████",
		"    █",
		"█████",
		"█    ",
		"█████",
	}

	three := placeholder{
		"█████",
		"    █",
		"█████",
		"    █",
		"█████",
	}

	four := placeholder{
		"█   █",
		"█   █",
		"█████",
		"    █",
		"    █",
	}

	five := placeholder{
		"█████",
		"█    ",
		"█████",
		"    █",
		"█████",
	}

	six := placeholder{
		"████ ",
		"█    ",
		"█████",
		"█   █",
		"█████",
	}

	seven := placeholder{
		"█████",
		"    █",
		"    █",
		"    █",
		"    █",
	}

	eight := placeholder{
		"█████",
		"█   █",
		"█████",
		"█   █",
		"█████",
	}

	nine := placeholder{
		"█████",
		"█   █",
		"█████",
		"    █",
		"    █",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	return digits
}

func getColon() placeholder {
	colon := placeholder{
		"     ",
		"  █  ",
		"     ",
		"  █  ",
		"     ",
	}
	return colon
}
