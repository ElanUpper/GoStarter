package main

import (
	"fmt"
	"time"
)

// sleep * n seconds
func sleeper(n int) {
	time.Sleep(time.Second)
}

// worker
func worker(ch <-chan rune, switcher int) {
	switch switcher {
	case 1: // for approach to handle channel
		for {
			if r_data, ok := <-ch; ok {
				fmt.Printf("receive %c . \n", r_data)
			} else { // 如果channel关闭
				fmt.Println("the channel is close.")
				return
			}

		}
	case 2: // range approach to handle channel
		for item := range ch {
			fmt.Printf("receive %c . \n", item)
		}

	}

}

// channel 只读/只写
func chann_demo() {

	ch := make(chan int)

	go func(ch <-chan int) {
		for {
			dt := <-ch
			fmt.Printf("receive %v\n", dt)
		}
	}(ch)

	time.Sleep(1 * time.Second)

	go func(ch chan<- int) {
		for i := 0; i < 20; i++ {
			ch <- i
		}
	}(ch)

}

// 说明概念
func make_array_channel_simple() {

	// 仅仅 chan_lt := [1]chan int{}
	// 那么chan_lt就为nil
	chan_lt := [1]chan int{make(chan int)}

	go func() {
		fmt.Println("send begin")
		chan_lt[0] <- 10
		fmt.Println("send end")
	}()

	go func() {
		fmt.Println("receive begin")
		fmt.Print(<-chan_lt[0])
		fmt.Println("receive end")
	}()

	time.Sleep(2 * time.Second)
}

// 具体实战
func make_array_channel_complex() {

	// list_ch := [10]chan int{}
	var list_ch [10]chan int

	// receive
	for i := 0; i < 10; i++ {
		list_ch[i] = make(chan int)
		go func(i int) {
			for {
				rec_data := <-list_ch[i]
				fmt.Printf("%d receive %c\n", i, rec_data)
			}
		}(i)
	}

	// send
	for i := 0; i < 10; i++ {
		go func(i int) {
			list_ch[i] <- 'a' + i
		}(i)

		go func(i int) {
			list_ch[i] <- 'A' + i
		}(i)
	}

}

//  创建一个消费者(chan作为返回值)
func CreateWorker() chan<- int {
	ch := make(chan int)
	go func() {
		for {
			fmt.Println("receive ", <-ch)
		}
	}()
	return ch
}

func Paramter_channel() {

	worker := CreateWorker()
	for i := 0; i < 10; i++ {
		fmt.Printf("start: %d send.\n", i)
		worker <- i
		fmt.Printf("end: %d send.\n", i)
	}

}

// 带缓存区的channel
func buffer_channel() {
	ch := make(chan rune, 2)
	//go worker(ch)
	ch <- 'a'
	ch <- 'b'
	//ch <- 'c';
}

func close_channel() {
	ch := make(chan rune)
	go worker(ch, 2)
	ch <- 'a'
	close(ch)
}

func main() {

	close_channel()

	sleeper(2)

}
