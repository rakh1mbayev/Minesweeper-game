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

	greeting()

	var x, y int
	var first_touch, win bool = true, true
	total_cell := n * m
	for moves_count+bombs_count < total_cell {

		x, y = inputData(true)

		x--
		y--

		if us[x][y] {
			fmt.Println("Error: you are already selected this cell.")
			continue
		}

		moves_count++
		if first_touch {
			if matrix[x][y] == '*' {
				replace(x, y)
			}
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
