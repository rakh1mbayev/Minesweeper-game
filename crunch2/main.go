package main

import (
	"fmt"
)

var (
	n, m, bombs_count, moves_count int
	matrix                         [][]rune
	postmatrix                     [][]int
	us                             [][]bool
)

func main() {

	// input data
	fmt.Scanln(&n, &m)

	createField()

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanln(&s)
		for j := 0; j < m; j++ {
			matrix[i][j] = rune(s[j])
		}
	}

	var x, y int
	var first_touch bool = true
	win := true
	for moves_count+bombs_count < n*m {
		OutPutting()

		fmt.Print("Enter coordinates: ")
		fmt.Scan(&x)
		fmt.Print(" ")
		fmt.Scan(&y)
		moves_count++

		if first_touch {
			first_touch = false
			calculate()
		}

		if postmatrix[x-1][y-1] == -1 {
			us[x-1][y-1] = true
			win = false
			break
		}
		if postmatrix[x-1][y-1] == 0 {
			blankTravel(x-1, y-1)
		}
		us[x-1][y-1] = true
	}
	OutPutting()
	statistics(win)
}
