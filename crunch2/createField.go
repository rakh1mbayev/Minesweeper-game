package main

func createField() {

	matrix = make([][]rune, n)
	postmatrix = make([][]int, n)
	us = make([][]bool, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]rune, m)
		postmatrix[i] = make([]int, m)
		us[i] = make([]bool, m)
	}
}
