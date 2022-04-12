package main

import "fmt"

func main() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(sum)

	sub := func(num1, num2 int) int {
		return num1 - num2
	}

	result := sub(10, 7)
	fmt.Println(result)
}
