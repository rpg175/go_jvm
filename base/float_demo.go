package main

import "fmt"

func main() {
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	//科学计数法，E大小写都可以，gofmt自动转为小写
	var num3 float32 = 314e-2
	fmt.Println(num3)
	var num4 float32 = 314e+2
	fmt.Println(num4)
	var num5 float32 = 314e+2
	fmt.Println(num5)
	//golang中默认的浮点类型为float64
	var num6 float64 = 314e+2
	fmt.Println(num6)

	//float32会丢失精度
	var num7 float32 = 3.00000014159
	fmt.Println(num7)
	var num8 float64 = 3.00000014159
	fmt.Println(num8)

}
