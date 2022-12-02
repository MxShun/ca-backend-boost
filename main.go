package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1 := scanNumber(1)
	num2 := scanNumber(2)

	fmt.Print(num1 + num2)
}

func scanNumber(index int) int {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input", index, ":")
	scanner.Scan()
	strNum := scanner.Text()
	num, _ := strconv.Atoi(strNum)

	return num
}