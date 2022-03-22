package main

import "fmt"

func testF(num int) {
	fmt.Println(num)
}

func main() {
	a := testF
	fmt.Printf("a %T, test %T \n", a, testF)
	a(10)

}
