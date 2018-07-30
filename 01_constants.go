package main

import . "fmt"  // 直接导入 如果使用fmt中的func不需要添加fmt.
import f "fmt"  // 直接导入fmt但是重命名为f

func const_define(){
	// const
	//   后面不使用不会报错
	//   可以不定义类型
	const numa, numb = 10, 20
	// 批量定义
	const (
		char_a = "a"
		numb_b = 30
		hex_c  = "x16"
	)
	f.Println(char_a, numb_b, hex_c)

}

// 定义枚举类型
func enum_define(){
	const (
		Cplusplus 		= 0
		GoLang 			= 1
		python		    = 2
	)
	const (
		Cplusplus_1 		= iota  // 保证自增 0起始
		GoLang_1
		_    				// 跳过一个计数值
		python_1   			// 3
	)
	const (
		Cplusplus_2 		= iota * 10 + 3 // 保证自增 0起始
		GoLang_2
		_
		_ // 跳过2个计数值
		tmp					// 43
		python_2   		    = iota << 2 + 2	//  iota=5(左移两位)=>101(00)=>20 + 2 =>22
		tmp1				// 110(00)+2 => 24 + 2 => 26

	)
	Println(Cplusplus_2, python_2, tmp1)
}

func main(){
	const_define()
	enum_define()
}
