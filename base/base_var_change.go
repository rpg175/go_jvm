package main

import "fmt"

func main() {
	var n1 int = 100
	//var n2 float64 = n1 报错
	var n2 float64 = float64(n1)
	fmt.Println(n1)
	fmt.Println(n2)

	//将int32转成int8，不会报错，但会损失精度
	var n3 int32 = 230001
	var n4 = int8(n3)
	fmt.Println(n3)
	fmt.Println(n4)

	var n5 int32 = 230001
	var n6 = int64(n5) + 30
	var n7 = int64(n5 + 30) //表达式左右的数据类型需要匹配
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)

	var n8 int32 = 126
	var n9 = int8(n8) + 1
	//var n10 = int8(n8) + 128 //表达式左右的数据类型需要匹配，128已经超过int8的表示范围，编辑器直接报错
	var n10 = int8(n8) + 2 //表达式左右的数据类型需要匹配
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)
}
