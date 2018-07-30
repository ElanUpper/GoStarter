package main

import (
	. "awesomeProject/scr/04_Fib"
	"bufio"
	"fmt"
	"os"
)

// defer 练习
func defer_try(filename string) {
	f, err := os.Create(filename)
	defer f.Close() // 确保最后可以关闭文件
	if err == nil {
		writer := bufio.NewWriter(f)
		// 创建fib生成器函数
		fib := Fib_gen()
		// 写入文件
		for i := 0; i < 20; i++ {
			fmt.Fprintln(writer, fib())
		}
		defer writer.Flush()
	}
}

// error 练习
func error_try(filename string){
	_, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 假如自定义err  测试！！
	// err = errors.New("custom error!! ")

	if err != nil {
		if Perr, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s\n", Perr.Op, Perr.Path, Perr.Err)
		} else {
			panic(err)  // 如果是自定义的就会走这个语句  
		}

	} else {
		fmt.Println("open successful")
	}

}

func main() {
	defer_try("writer.txt")
	error_try("writer.txt")
}
