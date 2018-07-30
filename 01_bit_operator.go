package main

import "fmt"

func main(){
	// 一元运算符取反: ^ !
	numb_a := 1    // 1010
	flag_a := true
	// ! 仅仅可以用到布尔
	fmt.Println(^numb_a,  // (0)1010 -> (1)0101
				!flag_a)  //

	// 放置除0异常
    var div1 int = 0 ;
    if div1 != 0 && ( 100 / div1 ) > 10 {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}

}