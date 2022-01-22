package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = &sync.WaitGroup{}
var mutex = &sync.Mutex{}
var n = 0

func testAddN() {
	n++
	fmt.Println(n)
}

func lockTestAddN() {
	mutex.Lock()
	defer mutex.Unlock()

	n++
	fmt.Println(n)
	time.Sleep(time.Second * 1)
}

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

	//times := 5
	//waitGroup.Add(times)
	//for i := 0; i < times; i++ {
	//	go func() {
	//		fmt.Println("hello,go func " + strconv.Itoa(i))
	//		waitGroup.Done()
	//	}()
	//}
	//waitGroup.Wait()

	//for i := 0; i < 5; i++ {
	//	go testAddN()
	//}
	//time.Sleep(time.Second * 5)
	//fmt.Println(n)

	//for i := 0; i < 5; i++ {
	//	go lockTestAddN()
	//}
	//time.Sleep(time.Second * 8)
	//fmt.Println(n)
}
