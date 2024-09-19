package main

import "fmt"

func main() {
	var i int
	fmt.Print("Введите число i: ")
	fmt.Scan(&i)
	switch {
	case i > 0:
		fmt.Println("i больше нуля")
	case i < 0:
		fmt.Println("i меньше нуля")
	default:
		fmt.Println("i равно нулю")
	}
}