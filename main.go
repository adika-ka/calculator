package main

import (
	"fmt"
	"strconv"
)

// Мапы чисел
type m = map[string]int
type aTr = map[string]string

// Переменные для принимаемых данных
var num1Str, operator, num2Str, over string

func main(){
	//Начало работы приложения, прием данных
	getArithmeticOperation()

	romanToArabic := make(m)

	romanToArabic["I"] = 1
	romanToArabic["II"] = 2
	romanToArabic["III"] = 3
	romanToArabic["IV"] = 4
	romanToArabic["V"] = 5
	romanToArabic["VI"] = 6
	romanToArabic["VII"] = 7
	romanToArabic["VIII"] = 8
	romanToArabic["IX"] = 9
	romanToArabic["X"] = 10
	romanToArabic["XX"] = 20
	romanToArabic["XXX"] = 30
	romanToArabic["XL"] = 40
	romanToArabic["L"] = 50
	romanToArabic["LX"] = 60
	romanToArabic["LXX"] = 70
	romanToArabic["LXXX"] = 80
	romanToArabic["XC"] = 90
	romanToArabic["C"] = 100

	arabicToRoman := make(aTr)

	arabicToRoman ["1"] = "I"
	arabicToRoman ["2"] = "II"
	arabicToRoman ["3"] = "III"
	arabicToRoman ["4"] = "IV"
	arabicToRoman ["5"] = "V"
	arabicToRoman ["6"] = "VI"
	arabicToRoman ["7"] = "VII"
	arabicToRoman ["8"] = "VIII"
	arabicToRoman ["9"] = "IX"
	arabicToRoman ["10"] = "X"
	arabicToRoman ["20"] = "XX"
	arabicToRoman ["30"] = "XXX"
	arabicToRoman ["40"] = "XL"
	arabicToRoman ["50"] = "L"
	arabicToRoman ["60"] = "LX"
	arabicToRoman ["70"] = "LXX"
	arabicToRoman ["80"] = "LXXX"
	arabicToRoman ["90"] = "XC"
	arabicToRoman ["100"] ="C"

	isRoman := false
	
	// Преобразование
	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if ((num1 == 0 && err2 == nil) || (num2 == 0 && err1 == nil))  {
		panic("Используются одновременно разные системы счисления")
	}
	//если ошибка
	if err1 != nil{
		num1 = romanToArabic[num1Str]
		if num1 != 0 {
			isRoman = true
		}
	}
	if err2 != nil{
		num2 = romanToArabic[num2Str]
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10{
		panic("Неккоректное число")
	}
	// АрОперации
	var result int
	switch operator{
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0{
			fmt.Println("Деление на ноль")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Неверный оператор")
		return
	}

	if isRoman {
		if result <= 0 {
			panic("В римской системе нет отрицательных чисел")
		}
		converting(result, arabicToRoman)
		return
	} 

	fmt.Println(result) 
}

//Функция для преобразования арабских чисел в римские
func converting (res int, arabicToRoman aTr){
	intToStr := strconv.Itoa(res)
	var result, rem, romRem string
	var remResult, arabRem int

	for key, value := range arabicToRoman{
		if intToStr == key {
			result = value
			fmt.Println(result)
			return
		} else if len(intToStr) == 2{
			rem = string(intToStr[1])
		}
	}
	for key, value := range arabicToRoman{
		if rem == key{
			romRem = value
		}
	}
	arabRem, _ = strconv.Atoi(rem)
	remResult = res - arabRem
	for key, value := range arabicToRoman{
		if strconv.Itoa(remResult) == key {
			result = value + romRem
			fmt.Println(result)
		}
	}

}

// Принимаем операцию и если не корректные данные паника
func getArithmeticOperation(){
	fmt.Println("Введите операцию (например, 1 + 2 или I + II): ")
	fmt.Scanf("%s %s %s %s", &num1Str, &operator, &num2Str, &over)

	if over != "" {
		panic("Не соответствует возможной арифметической операций")
	}
}





