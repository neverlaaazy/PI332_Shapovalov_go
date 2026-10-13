package main

import (
	"fmt"
	"strings"
)

const usd float64 = 83.87
const euro float64 = 93.84

func convertRuToUsd(ruMoney float64) float64 {
	var resultUSD float64 = ruMoney / usd
	return resultUSD
}
func convertRuToEuro(ruMoney float64) float64 {
	var resultEuro float64 = ruMoney / euro
	return resultEuro
}

func main() {
	var flag = false
	for !flag {
		var userMoneyRu float64
		var flagType = false
		for !flagType {
			fmt.Print("Введите сумму для конвертации в рублях: ")
			fmt.Scan(&userMoneyRu)
			if userMoneyRu > 0 {
				flagType = true
			} else {
				fmt.Sprint("err. Неверное значение, попробуйте ещё раз")
			}
		}

		fmt.Printf("Ваши рубли в долларах: %.2f usd,"+
			"\nВаши рубли в евро: %.2f euro.", convertRuToUsd(userMoneyRu), convertRuToEuro(userMoneyRu))

		var flagToExit = false
		for !flagToExit {
			fmt.Print("\nContinue(y/n): ")

			var userAnswer string
			fmt.Scan(&userAnswer)

			var userAnswerFL string = strings.ToLower(string(userAnswer[0]))
			if userAnswerFL == "n" {
				flag = true
				flagToExit = true
			} else if userAnswerFL == "y" {
				flagToExit = true
			} else {
				fmt.Println("err. Попробуйте ещё раз")
			}
		}

	}
}
