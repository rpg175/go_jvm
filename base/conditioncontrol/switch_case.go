package main

import "fmt"

func main() {
	var money int = 1000000

	//相对于其他语言，少了break
	switch money / 10000 {
	case 10000:
		fmt.Println("你是土豪")
	case 1000:
		fmt.Println("你小康之家")
	case 100:
		fmt.Println("温饱")
	case 10:
		fmt.Println("贫穷")
	default:
		fmt.Println("乞丐")
	}
}
