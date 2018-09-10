package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 函数一等公民（闭包）  不能有状态改变 =
// 参数，变量，返回值都可以是函数

func add_numb(i_type string) func(i_v int) (int, *int) {
	switch i_type {
	case "add":
		sum := 0
		// 返回sum的值 以及地址
		return func(v int) (int, *int) {
			sum += v
			return sum, &sum
		}
	default:
		return nil
	}

}

// 正统函数编程
type iAdder func(int) (int, iAdder)

func add_numb_func(i_a int) iAdder {
	return func(i_b int) (i_result int, ifunc iAdder) {
		return i_a + i_b, add_numb_func(i_a + i_b)
	}
}

// 斐波那契数列
func Fib_arry() func() int {
	i_a, i_b := 0, 1
	return func() int {
		i_a, i_b = i_b, i_a+i_b
		return i_a
	}
}

// 实现reader 接口
// 费布拉奇数列
type iReader func() int

func (r iReader) Read(p []byte) (n int, err error) {
	next := r()
	str := fmt.Sprintf("%d\n", next)
	if next > 20000 {
		return 0, io.EOF
	}
	return strings.NewReader(str).Read(p)
}

func printFibContents(reader iReader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	adder := add_numb("add")
	arr_nub := []int{2, 4, 8, 10}
	for i := 0; i < len(arr_nub); i++ {
		fmt.Println(adder(arr_nub[i]))
	}

	adder_func := add_numb_func(0)
	i_sum := 0
	for i := 1; i < 10; i++ {
		i_sum, adder_func = adder_func(i)
		fmt.Printf("1 + .... + %d = %d\n", i, i_sum)
	}

	funcb := Fib_arry()
	for i := 0; i < 20; i++ {
		fmt.Println(funcb())
	}

	printFibContents(funcb)

}
