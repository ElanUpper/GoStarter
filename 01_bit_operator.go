package main

import "fmt"

/*

^   !     					一元运算符
*   /  %  <<  >>  &  &^
+   -  |  ^                 二元运算符
==  != <  <=  >=  >
<-                          专门用于channel
&&
||


*/

func main(){
	// 一元运算符取反: ^ !
	numb_a := 10    // 1010
	flag_a := true
	// ! 仅仅可以用到布尔
	fmt.Println(^numb_a,  //
				!flag_a)  //

	// <<  等于乘以2*3
	fmt.Println(numb_a<<3)
	// <<  等于除以2*3 然后取整
	fmt.Println(numb_a>>3)
	// a&^b  将a&b首先二进制化  如果b中的位数中有1那么要求a中必须为0
	numb_b := 7  // 0111
	// numb_a &^ numb_b => 1010 &^ 0111 => 1000(8)
	fmt.Println(numb_a&^numb_b)

	// 放置除0异常
    var div1 int = 0 ;
    if div1 != 0 && ( 100 / div1 ) > 10 {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}

}