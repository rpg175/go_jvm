package main

import "fmt"

// 获取用户键盘输入语句
func main() {
	var age int
	fmt.Println("请录入学生的年龄:")
	fmt.Scanln(&age) //传入内存地址

	var name string
	fmt.Println("请录入学生的姓名:")
	fmt.Scanln(&name) //传入内存地址

	var score float32
	fmt.Println("请录入学生的分数:")
	fmt.Scanln(&score) //传入内存地址

	var isVip bool
	fmt.Println("请录入学生是否为VIP:")
	fmt.Scanln(&isVip) //传入内存地址

	fmt.Printf("学生的年龄为：%v，姓名为：%v, 成绩为 %v, 是否为VIP %v", age, name, score, isVip)
	fmt.Println("")
	fmt.Println(" 方式二................................")
	fmt.Println("请录入学生的年龄,姓名，分数，是否VIP，使用空格进行分割！")
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVip)
	fmt.Printf("学生的年龄为：%v，姓名为：%v, 成绩为 %v, 是否为VIP %v", age, name, score, isVip)
}
