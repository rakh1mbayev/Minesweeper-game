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

	createField() // initialization

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
		x--
		y--

		if first_touch {
			first_touch = false
			calculate()
		}
		us[x][y] = true
		if postmatrix[x][y] == -1 {
			win = false
			break
		}
		if postmatrix[x][y] == 0 {
			us[x][y] = false
			blankTravel(x, y)
		}
	}

	OutPutting()
	statistics(win)
}
