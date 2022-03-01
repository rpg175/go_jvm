package main

import "fmt"

func main() {
	//+加号
	//1.正数，2.相加操作 3.字符串拼接
	n1 := +10
	n2 := 4 + 7
	s1 := "abc" + "def"
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(s1)

	// /除号：
	fmt.Println(10 / 3)   //两int数据运算，结果是int
	fmt.Println(10.0 / 3) //两个float，结果是float

	// %取模 等价公式： a%b = (a - a/b*b)
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	//++自增操作，--自减操作
	var a int = 10
	a++ //加1操作
	fmt.Println(a)
	//go语言里，++，--操作非常简单，只能单独使用，不能参与到运算中去

}
