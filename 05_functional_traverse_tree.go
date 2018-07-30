package main

import "fmt"

type etree struct {
	val	 int ;
	left, right *etree ;
}

func (et *etree)tran_print(f func(*etree)) { // 传入参数中 定义函数的参数&返回值
	if et == nil {
		return ;
	}
	et.left.tran_print(f);
	f(et)
	et.right.tran_print(f)
}

// 1. 编写传入函数（注意参数以及返回值）
func printNode(et *etree){
	fmt.Println(et.val)
	et.val++ ;
}

// 2. 直接使用匿名函数（注意参数以及返回值）
func (node *etree) PrintTree() {
	node.tran_print(func(en *etree){
		fmt.Println(en.val)
	})
}

func main(){

	// init
	etree_1 := etree{10, nil, nil}
	etree_1.left = &etree{5, nil, nil}
	etree_1.left.left = &etree{3, nil, nil}
	etree_1.left.right = &etree{8, nil, nil}
	etree_1.right = &etree{15, nil, nil}
	etree_1.right.left = &etree{12, nil, nil}
	etree_1.right.right = &etree{20, nil, nil}
	// 1.传入刚才编写好的函数
	etree_1.tran_print(printNode)
	// 使用匿名函数
	etree_1.PrintTree()

}