package main

import "fmt"

func main() {
	//对应1 + 2 。。。。+100求和
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//for 循环实际就是让程序员写代码的效率高了，但底层该怎么执行还是怎么执行。底层效率没有变化

	var str string = "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	var str1 string = "hello world 你好"
	for i, value := range str1 {
		fmt.Printf("%d,%c \n", i, value)
	}
	//对str进行遍历，索引是i，值是value

	//双重循环,定义标签可以任意停止循环
label2:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i: %v, j: %v \n", i, j)
			if i == 2 && j == 2 {
				break label2
			}
		}
	}

}
