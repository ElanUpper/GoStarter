package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("%d cores. \n", runtime.NumCPU())
	child_count := 20

	// 解决子协程完全执行完成退出 (全局变量失败！！！！！)
	doneJob := make(chan bool)
	gv_count := 0
	for i := 0; i < child_count; i++ {
		go func(id int) {
			gv_count++
			fmt.Printf("current id: %d, count %d \n", id, gv_count)
			if gv_count >= child_count-1 {
				doneJob <- true
			}
		}(i)
	}
	<-doneJob

	// 带缓存的channel 等待子协程次数
	doneJobQue := make(chan int, child_count)
	for i := 0; i < child_count; i++ {
		go func(id int) {
			doneJobQue <- id
		}(i)
	}

	for i := 0; i < child_count; i++ { // 等待协程完成退出
		DoneId := <-doneJobQue
		fmt.Printf("id: %d \n", DoneId)
	}

	sGroup := sync.WaitGroup{}
	sGroup.Add(child_count)
	for i := 0; i < child_count; i++ {
		go func(id int) {
			sGroup.Done()
		}(i)
	}
	sGroup.Wait()

}
