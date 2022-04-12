package main

import (
	"fmt"
	"rpg175.com/jvm/base/functionDemo/utils"
)

var num int = testVar()

func testVar() int {
	fmt.Println("test 函数被执行了!")
	return 10
}

func init() {
	fmt.Println("init 函数被执行！")
}

func main() {
	fmt.Println("main 函数被执行!")
	fmt.Println(utils.Age)
	fmt.Println(utils.Name)
}
