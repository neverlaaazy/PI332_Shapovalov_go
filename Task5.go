package main

import (
	"fmt"
)

func IsLeapYear(number int) {

	if number%400 == 0 || (number%4 == 0 && number%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var number int
	var isCorrect = false
	for !isCorrect {
		fmt.Print("Введите любое положительное, целое и не > 10000: ")
		fmt.Scan(&number)

		if number <= 10000 && number > 0 {
			isCorrect = true
		} else {
			fmt.Println("Ошибка, год < 1 или > 10000")
		}
	}

	IsLeapYear(number)
}
