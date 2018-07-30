package main

import (
	"fmt"
	"strings"
)

func test_string_func(str string){
	fmt.Println(strings.Fields(str))
	str1 := strings.Split(str, " ")
	fmt.Println(str1)
}

func main(){
	test_string_func("hello world  test   erwer")
}
