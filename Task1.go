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
	WhatANumber(-1)
}
