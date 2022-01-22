package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup = &sync.WaitGroup{}

func say(text string) {
	fmt.Println(text)
	time.Sleep(time.Second * 1)
	waitGroup.Done()
}

func main() {
	//say("hello world")

	//go say("hello world")
	//time.Sleep(time.Second*5)

	//waitGroup.Add(1)
	//go say("hello world")
	//waitGroup.Wait()

	//times := 5
	//waitGroup.Add(times)
	//for i := 0; i < times; i++ {
	//	go say("hello," + strconv.Itoa(i))
	//}
	//waitGroup.Wait()

	times := 5
	waitGroup.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			fmt.Println("hello,go func " + strconv.Itoa(i))
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}
