package main

import "fmt"

func greeting() {
	fmt.Println("1. Enter a custom map")
	fmt.Println("2. Generate a random map")
	fmt.Print("Enter your choice: ")

	var ans string
	fmt.Scanf("%s\n", &ans)

	for ans != "1" && ans != "2" {
		fmt.Println("Error: invalid input.")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%s\n", &ans)
	}
	n, m = inputData(false)
	createField() // initialization
	if ans == "1" {
		customMap()
	} else {
		randomMap()
	}
}
