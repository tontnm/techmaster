package main

import (
	"fmt"
	"time"
)

func greeting() {
	currentHour := time.Now().Hour()
	switch {
	case currentHour < 10:
		fmt.Println("Good Morning")
	case currentHour < 16:
		fmt.Println("Good Afternoon")
	case currentHour < 24:
		fmt.Println("Good Evening")
	}
}
