package main

import (
	"fmt"
	"runtime"
)

var food chan string

func Consumer(idx int) {
	for {
		if v, ok := <-food; ok {
			fmt.Printf("index %d, 吃了 %s\n", idx, v)
			break
		}
	}
}

func Productor(idx int) {
	for {
		food_desc := ""
		if idx/2 == 0 {
			food_desc = "包子"
		} else {
			food_desc = "饺子"
		}
		fmt.Printf("index %d, 做了 %s\n", idx, food_desc)
		food <- food_desc
		break
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	loop_count := 30
	food = make(chan string, 3)
	for i := 0; i < loop_count; i++ {
		Productor(i)
		Consumer(i)
	}
}
