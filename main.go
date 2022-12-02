package main

import (
	"fmt"
)

func main() {
	n := 10
	exec1(n)
	n = 20
	exec2(&n)
	fmt.Printf("Finish : %d \n", n)
}

func exec1(num int) {
	fmt.Printf("exec1 : %d \n", num)
	num = 100
	fmt.Printf("exec1 : %d \n", num)
}

func exec2(num *int) {
	fmt.Printf("exec2 : %d \n", *num)
	*num = 200
	fmt.Printf("exec2 : %d \n", *num)
}
