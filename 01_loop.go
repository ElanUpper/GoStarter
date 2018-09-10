package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 2^n + 2^(n-1) + ... 2^0 = numb
// 所以先取模 算出低位
// 然后依次/2  分别算出各个位对应的数
func convertToBin(numb int) string {
	result := ""
	for ; numb > 0; numb /= 2 {
		numb_1 := numb % 2
		result = strconv.Itoa(numb_1) + result
	}
	return result
}

// for通过省略初始值&自增 达到while效果
func while_read_file(filename string) string {
	content := ""
	file_pointer, err := os.Open(filename)
	fmt.Println(filename)
	if err != nil { // 如果没有错误信息
		panic("open the file failed")
	} else {
		scanner := bufio.NewScanner(file_pointer)
		for scanner.Scan() {
			content += scanner.Text() + ", "
		}
	}
	return content
}

// for通过省略 初始值&判断截止条件&自举条件 完成死循环
func for_endless_loop() {
	count := 0
	for {
		fmt.Println("hello world")
		count++
		if count > 10 {
			break
		}
	}
}

func main() {
	fmt.Println(convertToBin(10))
	fmt.Println(while_read_file("C:\\Users\\Administrator\\go\\src\\awesomeProject\\scr\\books.txt"))
	for_endless_loop()
}
