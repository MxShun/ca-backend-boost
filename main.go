package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1 := scanNumber(1)

	ope := scanOperator()

	num2 := scanNumber(2)

	fmt.Print(calc(num1, num2, ope))
}

func scanNumber(index int) int {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input %d:", index)
	scanner.Scan()
	strNum := scanner.Text()
	num, err := strconv.Atoi(strNum)

	if (err != nil) {
		fmt.Println("正しい数値を入力してください")
		return scanNumber(index)
	}

	return num
}

func scanOperator() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("InputOperator", ":")
	scanner.Scan()
	str := scanner.Text()

	switch(str) {
	case "+", "-", "*", "/":
		return str
	default:
		fmt.Println("正しい演算子を入力してください")
		return scanOperator()
	}
}

func calc(num1, num2 int, ope string) int {
	switch(ope) {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}
