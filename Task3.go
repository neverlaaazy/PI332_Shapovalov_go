package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var isCorrect = false
	for !isCorrect {
		fmt.Print("Введите любое положительное и целое число: ")
		fmt.Scan(&number)

		if number <= 10000 {
			isCorrect = true
		} else {
			fmt.Println("Ошибка, число больше 10000")
		}
	}
	var strNumber = strconv.Itoa(number)
	var firstNumber = int(strNumber[0] - '0')

	fmt.Printf("Число: %v, Первая цифра: %v", number, firstNumber)
}
