package main

import (
	"fmt"
	"rpg175.com/jvm/base/packageDemo/crud/db"
)

func main() {
	db.GetConnect()
	fmt.Println("启动成功")
}
