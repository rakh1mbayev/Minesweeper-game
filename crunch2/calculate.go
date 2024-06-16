package main

func calculate() {
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if matrix[x][y] == '*' {
				bombs_count++
				postmatrix[x][y] = -1
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if validation(x+i, y+j) && postmatrix[x+i][y+j] != -1 {
							postmatrix[x+i][y+j]++
						}
					}
				}
			}
		}
	}
}
