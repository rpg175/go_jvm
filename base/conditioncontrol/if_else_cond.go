package main

import "fmt"

func main() {
	if money := 100; money < 1000 {
		fmt.Println("资金不足，需要上班")
	} else {
		fmt.Println("资金充足，可以躺平")
	}
}
