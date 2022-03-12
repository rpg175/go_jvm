package main

import "fmt"

func main() {
	var count int = 20
	if count < 30 {
		fmt.Printf("口罩不足，只剩 %v", count)
	}
	fmt.Println()
	if age := 36; age > 35 {
		fmt.Printf("中年危机出现了，年龄 %v", age)
	}
}
