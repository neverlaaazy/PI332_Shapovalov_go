package main

import "fmt"

func test() {
	var (
		i   int     = 1
		i8  int8    = 2
		i16 int16   = 3
		i32 int32   = 4
		i64 int64   = 5
		by  byte    = 6
		f32 float32 = 1.1
		f64 float64 = 1.2
		s   string  = "text"
		b   bool    = false
	)
	num := 1
	text := "text"
	flag := false

	fmt.Printf("i = %T; i8 = %T; i16 = %T; i32 = %T; i64 = %T;"+
		"\nby = %T;"+
		"\ns = %T;"+
		"\nb = %T;"+
		"\nnum = %T;"+
		"\ntext = %T;"+
		"\nflag = %T;"+
		"\nf32 = %T; f64 = %T.\n\n", i, i8, i16, i32, i64, by, s, b, num, text, flag, f32, f64)
}
func swap(s1 string, s2 string) {
	s3 := s1
	s1 = s2
	s2 = s3
}

func main() {
	//Задание 1.1
	var number1 int = 1
	var number2 = 2

	var byte1 byte = 1

	var numF1 float32 = 1.2
	var numF2 float64 = 2.1
	var numF3 = 3.2

	var text1 string = "text1"
	var text2 = "text2"

	var flag1 bool = false
	var flag2 = true
	fmt.Println("Задание 1.1:")
	fmt.Printf("number1 = %T; number2 = %T;"+
		"\ntext1 = %T; text2 = %T;"+
		"\nflag1 = %T; flag2 = %T;"+
		"\nnumF1 = %T; numF2 = %T; numF3 = %T;"+
		"\nbyte1 = %T.\n\n", number1, number2, text1, text2, flag1, flag2, numF1, numF2, numF3, byte1)
	test()

	//Задание 1.2
	fmt.Println("Задание 1.2:")
	var s1 string = "text1"
	var s2 string = "text2"

	fmt.Println("Изначальные данные: " + s1 + " " + s2)

	// swap(s1, s2)

	var s3 string = s1
	s1 = s2
	s2 = s3
	fmt.Println("Способ с 3-ей переменной: " + s1 + " " + s2)

	s1, s2 = s2, s1

	fmt.Println("Способ без 3-ей переменной: " + s1 + " " + s2)

}
