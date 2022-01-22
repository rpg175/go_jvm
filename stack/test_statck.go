package main

import "fmt"

func a() {
	fmt.Println("func a " + b())
}

func b() string {
	fmt.Println("func b -> " + c())
	return "func b()"
}

func c() string {
	fmt.Println("func c")
	return "func c()"
}

func main() {
	a()
}

//stack
//c
//b
//a
