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

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Vui lòng nhập chiều cao của bạn: ")

	str, _ := reader.ReadString('\n')

	str = strings.TrimSuffix(str, "\n")
	if height, err = strconv.ParseFloat(str, 64); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Vui lòng nhập cân nặng của bạn: ")
	str, _ = reader.ReadString('\n')

	str = strings.TrimSuffix(str, "\n")
	if weight, err = strconv.ParseFloat(str, 64); err != nil {
		fmt.Println(err.Error())
		return
	}

	bmi := CalculateBMI(height, weight)
	fmt.Printf("Chỉ số bmi của bạn là: %f", bmi)
}

func CalculateBMI(height, weight float64) float64 {
	return weight / (height * height)
}
