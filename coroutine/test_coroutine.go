package main

import (
	"fmt"
	"time"
)

func say(text string) {
	fmt.Printf(text)
}

func main() {
	go say("hello world")
	//say("hello world")
	time.Sleep(time.Second*5)
}

