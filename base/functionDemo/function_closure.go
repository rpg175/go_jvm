package main

import "fmt"

// 函数sum的函数返回值是一个函数
func sum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

func main() {
	sumF := sum()
	fmt.Println(sumF(1))
	fmt.Println(sumF(2))
}
