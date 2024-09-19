package main

import "fmt"

func main() {
	var pet string
	var age int

	fmt.Print("Какой у тебя питомец? ")
	fmt.Scan(&pet)

	fmt.Print("Сколько ему лет? ")
	fmt.Scan(&age)

	if age < 2 {
		small_age := fmt.Sprintf("У тебя %s, которой (которому) %d год (лет). Он (она) еще маленький", pet, age)
		fmt.Print(small_age)
	} else {
		great_age := fmt.Sprintf("У тебя %s, которой (которому) %d. Он (она) такой (такая) большой (большая)!", pet, age)
		fmt.Print(great_age)
	}
}