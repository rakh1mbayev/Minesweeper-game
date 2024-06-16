package main

func blankTravel(x, y int) {
	if !validation(x, y) {
		return
	}
	if postmatrix[x][y] > 0 {
		us[x][y] = true
	}
	if postmatrix[x][y] != 0 || us[x][y] {
		return
	}
	us[x][y] = true
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			blankTravel(x+i, y+j)
		}
	}
}
