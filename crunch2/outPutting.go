package main

import "fmt"

func OutPutting() {
	fmt.Print("   ")
	for i := 0; i < m; i++ {
		fmt.Print("   ", i+1, "   ")
		if i+1 < m {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	fmt.Print("   ")
	ceiling := "_______"
	for i := 0; i < m; i++ {
		fmt.Print(ceiling)
		if i+1 < m {
			fmt.Print("_")
		}
	}
	fmt.Println()

	x := 0
	for i := 0; i < n*3; i++ {
		if i%3 == 0 {
			x++
		}
		if i%3 == 1 {
			fmt.Print(x)
		} else {
			fmt.Print(" ")
		}
		fmt.Print(" |")
		for j := 0; j < m; j++ {
			if us[x-1][j] {
				if i%3 == 2 {
					fmt.Print("_______|")
					continue
				}
				if i%3 == 1 && postmatrix[x-1][j] == -1 {
					fmt.Print("   *   |")
					continue
				}
				if i%3 == 1 && postmatrix[x-1][j] != 0 {
					fmt.Print("   ")
					colourText(rune(postmatrix[x-1][j] + '0'))
					fmt.Print("   |")
				} else {
					fmt.Print("       |")
				}
			} else {
				fmt.Print("xxxxxxx|")
			}
		}
		fmt.Println()
	}
}
