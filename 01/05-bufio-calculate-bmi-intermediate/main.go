package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err    error
		height float64
		weight float64
	)

	height, err = readFromKeyboard("Chiều cao: ")
	if err != nil {
		fmt.Println(err.Error())
	}

	weight, err = readFromKeyboard("Cân nặng:")
	if err != nil {
		fmt.Println(err.Error())
	}

	bmi := CalculateBMI(height, weight)
	fmt.Printf("Chỉ số bmi của bạn là: %f", bmi)
}

func readFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	str, _ := reader.ReadString('\n')

	str = strings.TrimSuffix(str, "\n")
	if result, err = strconv.ParseFloat(str, 64); err != nil {
		return 0.0, err
	}

	return result, nil
}

func CalculateBMI(height, weight float64) float64 {
	return weight / (height * height)
}
