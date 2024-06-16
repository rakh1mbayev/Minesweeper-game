package main

import "math/rand"

func randomMap() {
	minBombs := 2
	maxBombs := max(n*m/2, minBombs) // Maximum number of bombs is the total number of cells
	bombNum := rand.Intn(maxBombs-minBombs+1) + minBombs

	// Initialize the matrix with '.'
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = '.'
		}
	}

	// Place bombs randomly
	for k := 0; k < bombNum; k++ {
		for {
			i := rand.Intn(n)
			j := rand.Intn(m)
			if matrix[i][j] == '.' { // Check if the cell is empty
				matrix[i][j] = '*'
				break
			}
		}
	}
}
