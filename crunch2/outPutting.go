package main

import "fmt"

func ceiling() {
	fmt.Print("   ")
	for i := 0; i < m; i++ {
		fmt.Print("   ", i+1, "   ")
	}
	fmt.Println()

	fmt.Print("   ")
	roof := "_______"
	for i := 0; i < m; i++ {
		fmt.Print(roof)
		if i+1 < m {
			fmt.Print("_")
		}
	}
	fmt.Println()
}

func OutPutting() {

	ceiling()
	distn := n * 3       // total distance to last row
	distm := (m + 1) * 7 // total distance to last column
	x := 0
	y := 0
	for i := 0; i < distn; i++ {
		if i%3 == 0 {
			x++
		}
		y = 0
		for j := 0; j < distm; j++ {
			if j == 0 {
				if i%3 == 1 {
					fmt.Print(x)
				} else {
					fmt.Print(" ")
				}
				continue
			}
			if j == 1 {
				fmt.Print(" ")
				continue
			}
			if (j+6)%8 == 0 {
				fmt.Print("|")
				y++
				if y == m+1 {
					y = 1
				}
				continue
			}
			if us[x-1][y-1] {
				if i%3 == 2 {
					fmt.Print("_")
					continue
				}
				if i%3 == 1 && (j+2)%8 == 0 {
					if postmatrix[x-1][y-1] == 0 {
						fmt.Print(" ")
						continue
					}
					if postmatrix[x-1][y-1] == -1 {
						fmt.Print("*")
						continue
					}
					fmt.Print(postmatrix[x-1][y-1])
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
}
