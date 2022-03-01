package main

import "fmt"

func main() {
	fmt.Println(1 == 9) //false
	fmt.Println(1 != 9) //true
	fmt.Println(1 > 9)  //false
	fmt.Println(1 < 9)  //true
	fmt.Println(1 <= 9) //true 并不取等
	fmt.Println(1 >= 9) //false
	fmt.Println("----------------------------")
	//逻辑于
	//短路于：只要第一个数值//表达式结果是false
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println("----------------------------")
	//或逻辑
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println("----------------------------")
	//非逻辑
	fmt.Println(!true)
	fmt.Println(!false)
}
