package main

import "fmt"

func main() {
	var money int = 10000
	if money >= 10000_0000 {
		fmt.Println("土豪")
	} else if money >= 1000_0000 {
		fmt.Println("豪")
	} else if money >= 100_0000 {
		fmt.Println("小康")
	} else if money >= 10_0000 {
		fmt.Println("温饱")
	} else if money >= 1_0000 {
		fmt.Println("贫穷")
	} else if money < 1_0000 {
		fmt.Println("乞丐")
	}
}
