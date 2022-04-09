package main

import "fmt"

func main() {
	months := []string{"Oct", "Nov", "Dec", "Jan", "Feb", "Mar", "Jul", "Aug", "Sep", "Apr", "May", "Jun"}
	for _, month := range months {
		fmt.Printf("%d - %s - %s\n", monthIndex(month), month, quarterInYear(month))
	}
}

func monthIndex(month string) int {
	switch month {
	case "Jan":
		return 1
	case "Feb":
		return 2
	case "Mar":
		return 3
	case "Apr":
		return 4
	case "May":
		return 5
	case "Jun":
		return 6
	case "Jul":
		return 7
	case "Aug":
		return 8
	case "Sep":
		return 9
	case "Oct":
		return 10
	case "Nov":
		return 11
	default:
		return 12
	}
}

func quarterInYear(month string) string {
	switch month {
	case "Jan", "Feb", "Mar":
		return "1st quarter"
	case "Apr", "May", "Jun":
		return "2nd quarter"
	case "Jul", "Aug", "Sep":
		return "3rd quarter"
	default:
		return "4th quarter"
	}
}
