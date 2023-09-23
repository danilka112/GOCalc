package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanValues = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

func RomanToArabic(num string) int {

	for _, r := range num {
		_, err := romanValues[r]
		if !err {
			fmt.Println("Неверный ввод! Введите 2 операнда и оперцию, отделив их пробелами (x + y)")
			os.Exit(0)
		}
	}

	result := 0
	prevValue := 0

	numCheck, err := strconv.ParseFloat(num, 32)
	if err == nil {
		fmt.Printf("Введённый операнд %f не является целочисленным ", numCheck)
	}

	for i := len(num) - 1; i >= 0; i-- {
		currentValue := romanValues[rune(num[i])]
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}

	if result <= 10 {
		return result
	} else {
		fmt.Println("Введённые операнды должны иметь значения от 1 до 10!")
		os.Exit(0)
	}
	return 0
}

func ArabicToRoman(num int) string {
	if num <= 0 {
		return ""
	}

	romanNumeral := ""

	romanSymbols := []struct {
		Symbol string
		Value  int
	}{
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	for _, symbol := range romanSymbols {
		for num >= symbol.Value {
			romanNumeral += symbol.Symbol
			num -= symbol.Value
		}
	}

	return romanNumeral
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	arr := strings.Split(input, " ")
	if len(arr) > 3 || len(arr) < 3 {
		fmt.Println("Неверный ввод! Введите 2 операнда и оперцию, отделив их пробелами (x + y)")
		return
	}

	num1 := arr[0]
	num2 := arr[2]
	op := arr[1]
	var check1, check2 bool

	operand1, err := strconv.Atoi(num1)
	if err != nil {
		operand1 = RomanToArabic(num1)
		check1 = true
	}
	operand2, err := strconv.Atoi(num2)
	if err != nil {
		operand2 = RomanToArabic(num2)
		check2 = true
	}

	switch {
	case op == "+":
		operand1 += operand2
		break

	case op == "-":
		if (check1 == true && check2 == true) && operand1 > operand2 {
			operand1 -= operand2
		} else {
			if check1 != true && check2 != true && operand1 <= operand2 {
				operand1 -= operand2
			} else {
				fmt.Println("В Римской нотации запись отрицательного числа или 0 невозможна!")
				return
			}
		}
		break

	case op == "*":
		operand1 *= operand2
		break

	case op == "/":
		if operand2 > 0 {
			operand1 /= operand2
		} else {
			fmt.Println("Делить на 0 нельзя!")
			return
		}
		break
	default:
		fmt.Println("Неверный ввод оператора! Используйте +, -, *, /")
		return
	}

	if check1 == true && check2 == true {
		fmt.Println(ArabicToRoman(operand1))
	} else {
		if check1 != true && check2 != true {
			fmt.Println(operand1)
		} else {
			fmt.Println("Неверный ввод оба операнда должны быть в одной нотации!")
		}
	}
}
