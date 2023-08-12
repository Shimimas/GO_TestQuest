package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

func roman_check(number string) bool {
	if number == "I" {
		return true
	} else if number == "II" {
		return true
	} else if number == "III" {
		return true
	} else if number == "IV" {
		return true
	} else if number == "V" {
		return true
	} else if number == "VI" {
		return true
	} else if number == "VII" {
		return true
	} else if number == "VIII" {
		return true
	} else if number == "IX" {
		return true
	} else if number == "X" {
		return true
	}
	return false
}

func arabian_check(number string) bool {
	if number == "1" {
		return true
	} else if number == "2" {
		return true
	} else if number == "3" {
		return true
	} else if number == "4" {
		return true
	} else if number == "5" {
		return true
	} else if number == "6" {
		return true
	} else if number == "7" {
		return true
	} else if number == "8" {
		return true
	} else if number == "9" {
		return true
	} else if number == "10" {
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

func validator(input_string string) bool {    
	s := strings.Split(input_string, " ")
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
    if validator(input_string) {
		fmt.Println("EXECUTABLE")
	}
}