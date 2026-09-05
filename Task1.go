package main

import "fmt"

func WhatANumber(number int) {
	if number > 0 {
		fmt.Println("Число положительное")
	} else if number < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}

func main() {
	var number int
	fmt.Print("Введите целое число:")
	fmt.Scan(&number)
	fmt.Printf("Ваше число: %d!\nИтого: ", number)
	WhatANumber(number)
}
