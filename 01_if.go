package main

import (
	"io/ioutil"
	"fmt"
	"math"
)

func if_block(){
	const lf_name = "C:\\Users\\elan\\go\\src\\goproj\\src\\raw.txt"
	//println(filename)
	content, err := ioutil.ReadFile(lf_name)
	if err != nil {
		fmt.Printf("reason: %s \n", err)
	} else {
		fmt.Printf("the value: %s \n", string(content))
	}

	// 等价于  一些计算 -> 用于判断结果
	if content, err = ioutil.ReadFile(lf_name); err != nil {
		fmt.Printf("reason: %s \n", err)
	}else{
		fmt.Printf("the value: %s \n", string(content))
	}
}

func switch_block(num_a, num_b int, operator string) int {
	var result int ;
	switch operator {
	case "+":
		result = num_a + num_b ;
	case "-":
		result = num_a - num_b ;
	case "*":
		result = num_a * num_b ;
	case "/": fallthrough  // 默认switch是带break，需要需要需要添加fallthrough
	case "div":
		result = num_a / num_b ;
	}
	return result ;
}

// 精度比较大小
// p为用户自定义的比较精度，比如0.00001
const MIN_PRECISE = 0.00005
func IsEqual(f1, f2, p float64) (float64, bool) {
	diff := f1 - f2 ;
	return diff, math.Abs(diff) < p
}

func if_return(){
	if num_a := 10; num_a == 10 {
		fmt.Println("correct")
		return
	} else {
		fmt.Println("incorrect")
		return
	}
}

func main(){
	// if_block()
	fmt.Printf("the value %d \n", switch_block(10, 20, "+"))
	fmt.Printf("the value %d \n", switch_block(100, 20, "/"))
	fmt.Printf("the value %d \n", switch_block(100, 20, "div"))
	// 测试两个浮点数是否一样  对比精度
	fmt.Println( IsEqual(33.3333, float64(100.0/3), MIN_PRECISE) )
	if_return()
}

