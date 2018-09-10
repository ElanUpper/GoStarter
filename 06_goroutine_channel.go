package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var g_prog chan bool
var g_over_sign bool

func time_control_catcher(i_ch chan string) {
	for {
		select {
		case v, ok := <-i_ch:
			if ok {
				fmt.Println(v)
			}
		case <-time.After(time.Second * 3):
			fmt.Println("time up!!")
			g_prog <- true
			break
		}
	}
}

func select_test(i_ch_int chan int, i_ch_str chan string) {

	for {
		select {
		case lv_i_carrier, ok := <-i_ch_int:
			if ok {
				fmt.Printf("catch %d\n", lv_i_carrier)
			}
		case lv_str_carrier, ok := <-i_ch_str:
			if ok {
				fmt.Printf("catch %s\n", lv_str_carrier)
			}
		case <-time.After(3 * time.Second):
			fmt.Println("time is up!!")
			g_prog <- true
			break
		}
	}

}

func Hero_Legend(p1 chan int, p2 chan int) {
	player1_vol := 1000
	player2_vol := 1000
	for {
		select {

		case v1, ok := <-p1:
			if ok && !g_over_sign {
				player1_vol = player1_vol - v1
				fmt.Printf("player1 be hited, %d lost, current heath point %d\n", v1, player1_vol)
				if player1_vol <= 0 {
					g_over_sign = true
					fmt.Println("game over, player2 win the game\n")
					g_prog <- true
				}
			} /* else if !ok {
				fmt.Printf("player1 issue, %v\n", v1)
				g_over_sign = true ;
				g_prog <- true ;
				break ;
			} */
		case v2, ok := <-p2:
			if ok && !g_over_sign {
				player2_vol = player2_vol - v2
				fmt.Printf("player2 be hited, %d lost, current heath point %d\n", v2, player2_vol)
				if player2_vol <= 0 {
					g_over_sign = true
					fmt.Println("game over, player1 win the game\n")
					g_prog <- true
				}
			} else if !ok {
				fmt.Printf("player2 issue, %v\n", v2)
				g_over_sign = true
				g_prog <- true
				break
			}

		case <-time.After(10 * time.Second):
			fmt.Println("time is up!!!")
		}

	}

}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//  3秒自动结束channel！！！！
	//ch := make(chan string)
	g_prog = make(chan bool) //
	//
	//go time_control_catcher(ch)
	//
	//
	//for i := 0 ; i < 100 ; i++{
	//	ch <- "hello world" + string(i)
	//}
	//
	//<- g_prog
	//close(ch)
	//close(g_prog)

	/*  // select 处理多管道信息
	ch_int := make(chan int)
	ch_str := make(chan string)
	go select_test(ch_int, ch_str)
	ch_int <- 10
	ch_str <- "hello1"
	ch_int <- 20
	ch_str <- "hello2"
	<- g_prog
	close(ch_int)
	close(ch_str)
	*/

	// 运动员准备开始
	player1 := make(chan int, 2)
	player2 := make(chan int, 2) // 要设置缓存 否则会死锁！！！！
	go Hero_Legend(player1, player2)
	g_over_sign = false

	// 给与随机打击
	for {
		fmt.Printf("status: %v\n", g_over_sign)
		if g_over_sign {
			fmt.Println("congratulation!!!", time.Now())
			break
		} else {
			player1 <- rand.Intn(7) + rand.Intn(3) + 1
			player2 <- rand.Intn(7) + rand.Intn(3) + 1
		}

	}

	fmt.Println("-------------- pending start --------------", (time.Now()).Format("2006-01-02 15:04:05"))
	<-g_prog
	fmt.Println("-------------- pending end --------------", time.Now())
	close(player1)
	fmt.Println("player1 left the game", time.Now())
	close(player2)
	fmt.Println("player2 left the game", time.Now())
	close(g_prog)

}
