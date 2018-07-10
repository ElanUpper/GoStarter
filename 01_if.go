package main

import (
	"io/ioutil"
	"fmt"
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


func main(){
	// if_block()
	fmt.Printf("the value %d \n", switch_block(10, 20, "+"))
	fmt.Printf("the value %d \n", switch_block(100, 20, "/"))
	fmt.Printf("the value %d \n", switch_block(100, 20, "div"))
}

