package main

import (
	"fmt"
	"runtime"
)

// 协程属于非抢占式
func none_preemptive() {

	array_a := [10]int{}

	for i := 0; i < 10; i++ {
		fmt.Printf("the executor %d is start.\n", i)
		go func(i int) {
			for {
				array_a[i]++
				runtime.Gosched()
			}
		}(i)
		fmt.Printf("the executor %d is done.\n", i)
	}
	fmt.Println(array_a)
	// 检查变量冲突
	// go run -race 08_go_routine.go
}

//
func main() {
	none_preemptive()
}
