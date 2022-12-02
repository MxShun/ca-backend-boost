package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input X:")
	scanner.Scan()
	x := scanner.Text()

	fmt.Print(x)
}