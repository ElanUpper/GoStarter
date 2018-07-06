package main

import "fmt"

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

func main(){
	fmt.Println("hello world");
	variable()
}