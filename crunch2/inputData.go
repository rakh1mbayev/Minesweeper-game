package main

import "fmt"

func inputData(coordinates bool) (int, int) {
	var stringA, stringB string

	if coordinates {
		OutPutting()
		fmt.Print("Enter coordinates: ")
	}
	fmt.Scanf("%s %s\n", &stringA, &stringB)

	for !inputValid(&stringA) || !inputValid(&stringB) {
		fmt.Println("Error: invalid input.")
		if coordinates {
			OutPutting()
			fmt.Print("Enter coordinates: ")
		} else {
			fmt.Print("Enter dimensions: ")
		}
		fmt.Scanf("%s %s\n", &stringA, &stringB)
	}

	return convert(&stringA), convert(&stringB)
}
