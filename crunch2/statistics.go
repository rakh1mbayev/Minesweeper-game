package main

import "fmt"

func statistics(win bool) {
	if win {
		fmt.Println("You Win!")
	} else {
		fmt.Println("Game Over!")
	}

	fmt.Println("Your statistics:")
	fmt.Println("Field size: ", n, "x", m)
	fmt.Println("Number of bombs: ", bombs_count)
	fmt.Println("Number of moves: ", moves_count)
}
