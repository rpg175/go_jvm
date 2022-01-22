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

	//ch := make(chan int) //Java : ch = new BlockingQueue<Integer> (size = 1)
	//defer close(ch)
	//
	//ch <- 1   // Java : ch.put(1) 阻塞方式往里塞数
	//x := <-ch // Java : ch.take() 阻塞方式往外拿数
	//fmt.Println(x)
	//fatal error: all goroutines are asleep - deadlock!

	//ch := make(chan int) //Java : ch = new BlockingQueue<Integer> (size = 1)
	//defer close(ch)
	//
	////需要在不同的携程里面分别拿数和取数
	//go func() {
	//	ch <- 1 // Java : ch.put(1) 阻塞方式往里塞数
	//}()
	//
	//x := <-ch // Java : ch.take() 阻塞方式往外拿数
	//fmt.Println(x)
	////output:1

	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//defer func() {
	//	close(ch1)
	//	close(ch2)
	//}()
	//
	//go func() {
	//	for i := 0; ; i++ {
	//		ch1 <- i
	//		time.Sleep(time.Second)
	//	}
	//}()
	//
	//go func() {
	//	for i := 0; ; i-- {
	//		ch2 <- i
	//		time.Sleep(time.Second * 5)
	//	}
	//}()
	//
	//for {
	//	x := <-ch1
	//	y := <-ch2 // Java : ch.take() 阻塞方式往外拿数，所以导致sleep 5秒后执行
	//	fmt.Println(x)
	//	fmt.Println(y)
	//}

	ch1 := make(chan int)
	ch2 := make(chan int)
	defer func() {
		close(ch1)
		close(ch2)
	}()

	go func() {
		for i := 0; ; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; ; i-- {
			ch2 <- i
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		//select case不需要阻塞等待，谁有数据处理谁（信号量的方式）
		select {
		case x := <-ch1:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		}
	}

}
