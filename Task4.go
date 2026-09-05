package main

import (
	"fmt"
	"strconv"
)

func IsHappyTicket(number int) {
	var firN = int(strconv.Itoa(number)[0] - '0')
	var SecN = int(strconv.Itoa(number)[1] - '0')
	var TherN = int(strconv.Itoa(number)[2] - '0')
	var FourN = int(strconv.Itoa(number)[3] - '0')
	var FifN = int(strconv.Itoa(number)[4] - '0')
	var SixN = int(strconv.Itoa(number)[5] - '0')

	if firN+SecN+TherN == FourN+FifN+SixN {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var number int
	var isCorrect = false
	for !isCorrect {
		fmt.Print("Введите любое положительное, целое и 6-и значное число: ")
		fmt.Scan(&number)

		if number <= 999999 && number >= 100000 {
			isCorrect = true
		} else {
			fmt.Println("Ошибка, число < 100000 или > 999999")
		}
	}

	IsHappyTicket(number)
}
