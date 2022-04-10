package main

import "fmt"

func main() {
	cars1 := [3]string{"Toyota", "BMW", "Ford"}   // array
	cars2 := [...]string{"Toyota", "BMW", "Ford"} // array
	cars3 := []string{"Toyota", "BMW", "Ford"}    // slice

	fmt.Println("Index loop")
	for i := 0; i < len(cars1); i++ {
		fmt.Println(i, cars1[i])
	}

	fmt.Println("Range loop")
	for i, c := range cars2 {
		fmt.Println(i, c)
	}

	fmt.Println("Conditional loop")
	i := 0
	for i < len(cars3) {
		fmt.Println(i, cars3[i])
		i++
	}
}
