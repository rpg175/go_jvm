package main

import "fmt"

func exchangeNum(num1, num2 int) {
	var t int
	t = num1
	num1 = num2
	num2 = t
}

func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("before exchange : num1 = %v, num2 = %v \n", num1, num2)
	exchangeNum(num1, num2)
	fmt.Printf("after exchange : num1 = %v, num2 = %v", num1, num2)
}
