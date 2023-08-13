package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

var (
	romans = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	arabian = map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	}

	romans_for_convertation = []string {"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	decimals = []int {100, 90, 50, 40, 10, 9, 5, 4, 1}
)

func to_roman(num int) string {
    roman := ""
    for num > 0 {
        for i := 0; i < 9; i++ {
			for num >= decimals[i] {
				roman += romans_for_convertation[i]
				num -= decimals[i]
			}
		}
	}
    return roman
}

func roman_check(number string) bool {
	if romans[number] != 0 {
		return true
	}
	return false
}

func arabian_check(number string) bool {
	if arabian[number] != 0 {
		return true
	}
	return false
}

func operand_check(operand string) bool {
	if operand == "+" {
		return true
	} else if operand == "-" {
		return true
	} else if operand == "*" {
		return true
	} else if operand == "/" {
		return true
	}
	return false
}

func operations(n1 int, n2 int, operation string) int {
	if operation == "+" {
		return n1 + n2
	} else if operation == "-" {
		return n1 - n2
	} else if operation == "*" {
		return n1 * n2
	} else {
		return n1 / n2
	}
}

func execute(s []string) {
	if roman_check(s[0]) {
		result := operations(romans[s[0]], romans[s[2]], s[1])
		if result < 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		} else {
			fmt.Printf("%s\n", to_roman(result))
		}
	} else {
		fmt.Printf("%d\n", operations(arabian[s[0]], arabian[s[2]], s[1]))
	}
}

func validator(s []string) bool {    
	if len(s) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return false
	}
	if !((roman_check(s[0]) && roman_check(s[2])) || (arabian_check(s[0]) && arabian_check(s[2]))) {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return false
	}
	if !operand_check(s[1]) {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return false
	}
    return true
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input_string := scanner.Text()
	s := strings.Split(input_string, " ")
    if validator(s) {
		execute(s)
	}
}