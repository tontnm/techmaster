package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Nhập cạnh a: ")
	fmt.Scanf("%d", &a)
	checkSideNotNegativeNumber(a)

	fmt.Println("Nhập cạnh b: ")
	fmt.Scanf("%d", &b)
	checkSideNotNegativeNumber(b)

	fmt.Println("Nhập cạnh c: ")
	fmt.Scanf("%d", &c)
	checkSideNotNegativeNumber(c)

	checkTriangle(a, b, c)
}

func checkTriangle(a, b, c int) {

	if a+b > c && a+c > b && b+c > a {
		if a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b {
			fmt.Println("Tam giác vuông")
		} else if a == b && b == c {
			fmt.Println("Tam giác đều")
		} else if a == b || a == c || b == c {
			fmt.Println("Tam giác cân")
		} else if a == b || a == c || b == c {
			fmt.Println("Tam giác cân")
		} else if a*a > b*b+c*c || b*b > a*a+c*c || c*c > a*a+b*b {
			fmt.Println("Tam giác tù")
		} else {

			fmt.Println("Tam giác nhọn")
		}
	} else {
		fmt.Println("Ko phải cạnh của tam giác")
	}
}

func checkSideNotNegativeNumber(num int) {
	if num < 0 {
		fmt.Println("Ko phải cạnh tam giác")
		return
	}
}
