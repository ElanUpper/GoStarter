package main

import "fmt"

/*

go语言

  自举: 开源并且自己编写自己！

  静态类型&编译型:  变量必须在什么时候指定类型，且不可改变； 不过语法简洁类似解释形

  跨平台&垃圾自动回收

  原生并发: goroutine & channel

  完善的构建工具： Go 编译，测试，安装，运行，分析 工具齐全

  多编程范式： 函数式编程 -> 面向对象（组合代替继承）

*/

func variable(){
	var li_age  int
	var ls_name string
	// 一起初始化
	var ls_para1, ls_para2 string = "p=1", "p=2" // assign value while variable is declare
	var li_parm1, ls_parm2, lb_parm3 = 10, "string", true // assign value after declare the variable
	//var (
	//	li_parm1_s, ls_parm2_s, lb_parm3_s = 10, "string", true
	//)
	li_parm1_, ls_parm2_, lb_parm3_ := 10, "string", true // := instead of var
	ls_name = "elan";
	li_age  = 20 ;
	fmt.Printf("age: %d, name: \"%s\"\n", li_age, ls_name)
	// format
	/*
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	 */
	fmt.Printf("paramter1: %s, parameter2: %s\n", ls_para1, ls_para2)
	// after format the string
	fmt.Printf("paramter1: %d, parameter2: %s, parameter3: %t\n",
				li_parm1, ls_parm2, lb_parm3)
	// after format the string
	fmt.Printf("paramter1: %d, parameter2: %s, parameter3: %t\n",
		li_parm1_, ls_parm2_, lb_parm3_)

}

func constant_f(){
	const filename = "a.txt"
	const (
		lv_1, lv_2 = 10, "hello"
	)
	fmt.Println(filename, lv_1, lv_2)
	// 枚举类型
	const (
		cpp    = 0
		java   = 1
		python = 4
	)
	const (
		cpp1 = iota  // 枚举类型起始值0
		_  // 跳过一个数字
		java1
		___
		python1
	)
	fmt.Println(cpp, cpp1, python, python1)
	//
	const (
		numb1 = 10 << ( iota * 2 )
		numb2
		numb3
		numb3_1 = iota
		numb4 = 4 << ( iota * 4 )  // 4 * 2 ^ ( 4 * 4 )
		numb5
	)
	fmt.Println(numb1, numb2, numb3, numb3_1, numb4, numb5)
}


func main(){
	fmt.Println("hello world");
	variable()
	constant_f()
}