package main

import "fmt"

// 首先，我们需一个 echo()函数，
// 其会把一个整型数组放到一个Channel中，
// 并返回这个Channel
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _,n:=range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 平方函数
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n:= range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}

// 过滤奇数函数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n:= range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// 求和函数
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n:=range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

// 下面的代码类似于我们执行了Unix/Linux命令：
// echo $nums | sq | sum
func test1() {
	var nums = []int{1,2,3,4,5,6,7,8,9,10}
	for n:= range sum(sq(odd(echo(nums)))) {
		fmt.Println(n)
	}
}

type EchoFunc func([]int) (<- chan int)
type PipeFunc func(<- chan int) (<- chan int)

func pipeline(nums []int,echo EchoFunc, pipeFns ... PipeFunc) <- chan int {
	ch := echo(nums)
	for i:= range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func test2() {
	var nums = []int{1,2,3,4,5,6,7,8,9,10}
	for n:= range pipeline(nums,echo,odd,sq,sum) {
		fmt.Println(n)
	}
}

func main() {
	test1()
	test2()
}