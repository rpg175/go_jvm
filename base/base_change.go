package main

import "fmt"

func main() {
	//交换a b的值
	var a int = 8
	var b int = 4
	var t int = 0
	t = a
	a = b
	b = t
	fmt.Println(a, b)
}
