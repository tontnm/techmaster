package main

import "fmt"

func main() {
	var a, b, n int

	a = 0
	b = 1

	fmt.Printf("How many number? ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Fibonaci series are: ")
	fmt.Printf("%d %d ", a, b)

	printFibonacii(a, b, n-2)
}

func printFibonacii(a, b, n int) {
	var sum int

	if n > 0 {

		sum = a + b
		fmt.Printf("%d ", sum)

		a = b
		b = sum

		printFibonacii(a, b, n-1)
	}
}
