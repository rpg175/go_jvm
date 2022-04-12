package utils

import "fmt"

var Age int
var Name string

func init() {
	Age = 10
	Name = "我init了"
	fmt.Println("init 02 被执行了")
}
