package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string --> bool
	var s1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b type is %T, b=%v \n", b, b)

	//string --> int64
	var s2 string = "19"
	var num1 int64
	num1, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num1 type is %T, num1=%v \n", num1, num1)

	//string --> float64
	var s3 string = "3.14"
	var num2 float64
	num2, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("num2 type is %T, num2=%v \n", num2, num2)

	//string 转换基本类型，需要注意无法转换的时候，会回传默认值
	var s4 string = "golang"
	var b1 bool
	b1, _ = strconv.ParseBool(s4)
	fmt.Printf("b1 type is %T, b1=%v \n", b1, b1)

	var s5 string = "golang"
	var num3 int64
	num3, _ = strconv.ParseInt(s5, 10, 64)
	fmt.Printf("num3 type is %T, num3=%v \n", num3, num3)
}
