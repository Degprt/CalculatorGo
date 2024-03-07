package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	st := myScanner.Text()
	ex := foundTranslatingType(strings.Split(st, " "))
	fmt.Println(ex.Calculate())

}

func foundTranslatingType(arr []string) Expression {
	if len(arr) < 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
	} else if len(arr) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	result := "Arabian"
	num1, err1 := strconv.Atoi(arr[0])
	if err1 != nil {
		num1 = mapRoman[arr[0]]
		result = "Roman"
	}
	num2, err2 := strconv.Atoi(arr[2])
	if err2 != nil {
		num2 = mapRoman[arr[2]]
	}
	if err1 != nil && err2 == nil || err1 == nil && err2 != nil {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if num1 > 10 || num2 > 10 || num1 < 1 || num2 < 1 {
		panic("Выдача паники, так как операнды должны быть в диапозоне от 1 до 10")
	}
	ex := Expression{num1, arr[1], num2, result}
	return ex
}

type Expression struct {
	Num1    int
	Operand string
	Num2    int
	typeNum string
}

func (a *Expression) Calculate() string {
	var resultInt int
	switch a.Operand {
	case "+":
		resultInt = a.Num1 + a.Num2
	case "-":
		resultInt = a.Num1 - a.Num2
	case "*":
		resultInt = a.Num1 * a.Num2
	case "/":
		resultInt = a.Num1 / a.Num2
	default:
		panic("Выдача паники, так как оператор неизвестен (+, -, /, *).")
	}
	if a.typeNum == "Arabian" {
		return strconv.Itoa(resultInt)
	} else if a.typeNum == "Roman" && resultInt < 1 {
		panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
	} else {
		return arrRoman[resultInt]
	}
}

var mapRoman = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arrRoman = []string{
	"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII",
	"XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII",
	"XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII",
	"XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII",
	"XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV",
	"LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI",
	"LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV",
	"LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII",
	"LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI",
	"XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}
