package main

import "fmt"

func main() {
	testBMIs := []float64{14, 16, 16.5, 18, 23, 27, 32, 40, 42}
	for _, bmi := range testBMIs {
		fmt.Println(ReviewBMI(bmi))
	}
}

func ReviewBMI(bmi float64) string {
	switch {
	case bmi < 16:
		return "Severe thinness"
	case bmi < 16.9:
		return "Moderate thinness"
	case bmi < 18.4:
		return "Mild thinness"
	case bmi < 24.9:
		return "Normal"
	case bmi < 29.9:
		return "Pre-obese"
	case bmi < 34.9:
		return "Obese (Class I)"
	case bmi < 39.9:
		return "Obese (Class II)"
	default:
		return "Obese (Class III)"
	}
}
