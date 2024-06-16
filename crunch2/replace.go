package main

import "math/rand"

func replace(x, y int) {
	new_x := rand.Intn(n)
	new_y := rand.Intn(m)
	for matrix[new_x][new_y] == '*' {
		new_x = rand.Intn(n)
		new_y = rand.Intn(m)
	}
	matrix[x][y] = '.'
	matrix[new_x][new_y] = '*'
}
