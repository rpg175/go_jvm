package main

import "fmt"

func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main() {
	test()
	fmt.Println("-------")
	test(1, 2)
	fmt.Println("-------")
	test(1, 2, 3, 4, 5, 6)
}
