package main

import "fmt"

func main() {
	cars := []string{"Toyota", "BMW", "Ford"} // slice

	// recommend
	for i := len(cars) - 1; i >= 0; i-- {
		fmt.Println(i, cars[i])
	}

	// not recommend
	for i, c := range cars {
		defer fmt.Println(i, c)
	}
}
