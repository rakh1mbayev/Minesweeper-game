package main

import "fmt"

func colourText(char rune) {
	switch char {
	case '1':
		fmt.Printf("\x1b[31m%c\x1b[0m", char) // Red
	case '2':
		fmt.Printf("\x1b[32m%c\x1b[0m", char) // Green
	case '3':
		fmt.Printf("\x1b[33m%c\x1b[0m", char) // Yellow
	case '4':
		fmt.Printf("\x1b[34m%c\x1b[0m", char) // Blue
	case '5':
		fmt.Printf("\x1b[35m%c\x1b[0m", char) // Magenta
	case '6':
		fmt.Printf("\x1b[36m%c\x1b[0m", char) // Yellow
	case '7':
		fmt.Printf("\x1b[37m%c\x1b[0m", char) // Blue
	case '8':
		fmt.Printf("\x1b[38m%c\x1b[0m", char) // Magenta
	default:
		fmt.Printf("%c", char)
	}
}
