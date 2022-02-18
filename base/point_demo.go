package main

import "fmt"

func main() {
	var age int = 8
	//&符号+变量，可以获取变量的内存地址
	fmt.Println(&age)

	//定义一个指针变量 ptr
	//ptr 指针变量的名字
	//ptr 对应的类型是: *int是一个指针类型（可以理解为指向int类型的指针）
	//&age就是一个地址
	var ptr *int = &age
	fmt.Println(ptr)
	//需要获取ptr的具体数据，* + 变量名
	fmt.Printf("prt指向的数值为: %v \n", *ptr)
	//1. 可以通过指针改变指向值
	*ptr = 20
	fmt.Printf("age的数值为: %v \n", age)

}
