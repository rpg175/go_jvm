package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int
	age = 18
	fmt.Printf("age %v \n", age)

	var n1 int = 19
	var n2 float64 = 4.78
	var n3 bool = false
	var n4 byte = 'a'

	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("S1 %T, s1 = %v \n", s1, s1)

	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("S2 %T, s2 = %v \n", s2, s2)

	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("S3 %T, s3 = %v \n", s3, s3)

	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("S4 %T, s4 = %v \n", s4, s4)

	var s5 string = strconv.FormatInt(int64(n1), 10) //第二个参数指定字面值的进制形式为10进制
	fmt.Printf("S1 %T, s1 = %q \n", s5, s5)

	var s6 string = strconv.FormatFloat(float64(n2), 'f', 9, 64)
	fmt.Printf("s6 %T, s6 = %q \n", s6, s6)

	var s7 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s7 %T, s7 = %q \n", s7, s7)

	var s8 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s8 %T, s8 = %q \n", s8, s8)
}
