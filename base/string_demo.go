package main

import "fmt"

func main() {
	var s1 = "就是这条街最靓的崽"
	fmt.Println(s1)

	//字符串是不可以变的：指的是字符串一旦定义好，其中的字符的值不能改变 //s2[1] = 't'报错
	var s2 string = "就是这条街最靓的崽"
	s2 = "def"
	fmt.Println(s2)

	//字符串：
	//没有特殊字符，用""
	//有特殊字符，用``
	var s3 string = "就是这条街最靓的崽"
	fmt.Println(s3)
	var s4 string = `func main() {
					var age int
					age = 18
					fmt.Printf("age %v", age)
				   }`
	fmt.Println(s4)

	s5 := "abc" + "efg"
	s5 += "h"
	fmt.Println(s5) //abcefgh

	//可读性的测试
	s6 := "abc" + "efg" +
		"add" +
		"ccc"
	fmt.Println(s6)
}
