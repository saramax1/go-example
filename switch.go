package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("jomeeee")
	default:
		fmt.Println("baghie hafte")

	}

	fmt.Println(time.Now())
	fmt.Println(time.Now().Weekday())
	fmt.Printf("%T", time.Now().Hour())
}
