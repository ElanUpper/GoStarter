package main

import (
	. "./04_Fib"
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

func func_buffer_io(file_name string) {
	if f, err := os.Create(file_name); err == nil {
		defer f.Close()

		gener := FunClosure()

		fwriter := bufio.NewWriter(f)
		defer fwriter.Flush()

		for i := 0; i < 20; i++ {
			fmt.Fprintln(fwriter, gener())
		}
	}
}

// error 练习
func Error_try(filename string) {
	_, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 假如自定义err  测试！！
	// err = errors.New("custom error!! ")

	if err != nil {
		if Perr, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s\n", Perr.Op, Perr.Path, Perr.Err)
		} else {
			panic(err) // 如果是自定义的就会走这个语句
		}

	} else {
		fmt.Println("open successful")
	}

}

func defer_between_return_1() (t int) {
	defer func() {
		t = t + 1
	}()

	return t
}

func defer_between_return_2() (r int) {
	t := 0

	defer func() {
		t = t + 1
	}()

	return t
}

// 这里的t实际上是defer函数中的参数
func defer_between_return_3() (t int) {
	defer func(t int) {
		t = t + 1
	}(t)

	return 0
}

// 这里的t实际上是形式参数的值
func defer_between_return_4() (t int) {
	defer func() {
		t = t + 1
	}()

	return 0
}

func main() {
	//func_buffer_io("writer.txt")
	//func_buffer_io("writer.txt")
	fmt.Println(defer_between_return_1())
	fmt.Println(defer_between_return_2())
	fmt.Println(defer_between_return_3())
	fmt.Println(defer_between_return_4())
}
