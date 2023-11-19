package main

import (
	"fmt"
	"strings"
)

func main() {

	sumFunc := func(x, y int) int {
		return x + y
	}
	subFunc := func(x, y int) int {
		return x - y
	}
	multFunc := func(x, y int) int {
		return x * y
	}
	divFunc := func(x, y int) int {
		return x / y
	}

	calc := func(x, y int, operation func(x, y int) int) int {
		return operation(x, y)
	}

	fmt.Println(calc(3, 7, sumFunc))
	fmt.Println(calc(3, 7, subFunc))
	fmt.Println(calc(3, 7, multFunc))
	fmt.Println(calc(3, 7, divFunc))

	message := "I love strings"
	splittedstring := strings.Split(message, " ")
	fmt.Println(splittedstring[0])
	fmt.Println(splittedstring[1])
	fmt.Println(splittedstring[2])
}
