package main

import "fmt"

func customMap() {
	var s string
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &s)
		if !checkMap(&s) {
			fmt.Println("Error: invalid input. Try again")
			i--
			continue
		}
		for j := 0; j < m; j++ {
			matrix[i][j] = rune(s[j])
		}
	}
}
