package main

import "fmt"

func main() {
	weight := 68.0
	height := 1.65
	bmi := calculateBMI(weight, height)
	fmt.Printf("Chỉ số bmi của bạn là: %f", bmi)
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}
