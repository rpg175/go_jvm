package main

import (
	"fmt"
	"strconv"
	"strings"
)

//如果你问，500数字里面包含1的，有多少个
func main() {
	count := 0
	for i := 1; i < 500; i++ {
		if strings.Contains(strconv.FormatInt(int64(i), 10), "1") {
			count += 1
			fmt.Print(i)
			fmt.Print(",")
		}
	}
	fmt.Println()
	fmt.Printf("total: %v", count)
}
