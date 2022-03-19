package main

import "fmt"

//为什么要使用函数：减少重复代码，进行代码复用，提升可维护性
func cal(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	var num1 int = 10
	var num2 int = 20
	sum := cal(num1, num2)
	fmt.Printf("sum %v \n", sum)
}
