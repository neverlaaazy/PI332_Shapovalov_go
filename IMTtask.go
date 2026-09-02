package main

import "fmt"

func calculateIMT(weight float64, height float64) float64 {
	return weight / ((height / 100) * (height / 100))
}

func ValuesOfIndicatorsIMT(weight float64, height float64) string {
	var resultIMT = calculateIMT(weight, height)

	if resultIMT < 18.5 {
		return "недостаток массы тела"
	} else if resultIMT >= 18.5 && resultIMT <= 24.9 {
		return "норма"
	} else if resultIMT >= 25 && resultIMT <= 29.9 {
		return "избыточная масса тела"
	} else {
		return "ожирение"
	}

}

func PrintAllInfoOfIMT(weight float64, height float64) {
	var intIMT = calculateIMT(weight, height)
	var stringIMT = ValuesOfIndicatorsIMT(weight, height)

	fmt.Println(fmt.Sprintf("\nВес(кг) = %.2f, Рост(см) = %.2f.\nIMT = %.2f, Итог = %s.", weight, height, intIMT, stringIMT))
}

func main() {

	fmt.Println(calculateIMT(100, 190))
	fmt.Println(ValuesOfIndicatorsIMT(100, 190))
	PrintAllInfoOfIMT(100, 190)
}