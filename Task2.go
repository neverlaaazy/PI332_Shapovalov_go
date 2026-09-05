package main

import "fmt"

func CheckForDistinctDigits(number int) string {
	digit3 := number % 10
	digit2 := (number / 10) % 10
	digit1 := number / 100

	if digit1 == digit2 || digit2 == digit3 || digit1 == digit3 {
		return "NO"
	}

	return "YES"
}

func main() {
	var number int
	var isCorrect = false
	for !isCorrect {
		fmt.Print("Введите 3-значное число: ")
		fmt.Scan(&number)

		if number >= 100 && number <= 999 {
			isCorrect = true
		} else {
			fmt.Println("Ошибка! Число должно быть строго от 100 до 999.")
		}
	}

	fmt.Printf("Ваше число: %v!\nИтог: ", number)

	var resultOfFunc = CheckForDistinctDigits(number)

	fmt.Println(resultOfFunc)
	if resultOfFunc == "YES" {
		fmt.Printf("Ни одна цифра не равняется другой в %v", number)
	} else {
		fmt.Printf("Одна из цифр равняется другой в %v", number)
	}
}
