package main

import "fmt"

func main() {
	c1 := 'a'
	fmt.Println(c1) //97
	c2 := '6'
	fmt.Println(c2) //54
	c3 := '-'
	fmt.Println(c3) //45
	c4 := 'c'
	fmt.Println(c4) //99
	c5 := '中'
	fmt.Println(c5)         //20013
	fmt.Printf("%c \n", c5) //中

	//转义字符
	//\n 换行
	fmt.Println("aaa\nbbb")
	//\b 退格
	fmt.Println("aaa\bbbb")
	//\r 光标回到本行开头，后续输入就会替换原有的字符
	fmt.Println("aaaaa\rbbb")
	//\t 制表符
	fmt.Println("aaa\tbb")
	//\"
	fmt.Println("\"Golang\"")
	//\'
	fmt.Println("'Golang'")

}
