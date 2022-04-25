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

//闭包：返回的匿名函数+匿名函数以外的变量num
//感受：匿名函数中引用的那个变量是会一直保存在内存中，可以一直使用
func main() {
	sumF := sum()
	fmt.Println(sumF(1)) //1
	fmt.Println(sumF(2)) //3
	fmt.Println(sumF(3)) //6
}
