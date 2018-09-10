package main

import "fmt"

func upd_slice(arr []int) {

	for k, v := range arr {
		fmt.Printf("index %d value %d\n", k, v)
	}
	// 直接更新数字
	arr[0] = 1000
}

// 跨域访问数据
func access_slice(arr []int) {
	// 打印全集
	fmt.Printf("original array: %v, len(arr): %d, cap(arr): %d\n", arr[:], len(arr), cap(arr))
	// 打印1-3
	fmt.Println("array[1-2]: ", arr[0:2])
	// 打印2-7
	fmt.Println("array[2-6]:: ", arr[1:6])
	// 强制直接越界 会发现后面自动补充0
	fmt.Println("array[2-10]:: ", arr[1:10])

}

// 初始化slice
func init_slice() {
	var slice1 []int
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i*2+1)
		// 我们会发现cap是 2^n进行扩增
		fmt.Printf("i:%d len:%d cap:%d\n", i, len(slice1), cap(slice1))
	}
	// 直接赋值
	slice2 := []int{2, 3, 4}
	fmt.Println(slice2)
	// 定义长度为15
	slice3 := make([]int, 15)
	slice3_1 := [15]int{}
	fmt.Printf("slice3 len:%d cap:%d\n", len(slice3), cap(slice3))
	fmt.Printf("slice3_1 len:%d cap:%d\n", len(slice3_1), cap(slice3_1))
	// 定义长度为10 capacity为18
	slice4 := make([]int, 10, 24)
	fmt.Printf("slice4 len:%d cap:%d\n", len(slice4), cap(slice4))
}

func manipulate_slice(arr []int) {
	tmp1 := make([]int, 20) // 一定要有length  如果修改为tmp1 := []int{} 那么将无法copy
	// 复制
	fmt.Printf("the original list %v\n", arr)
	copy(tmp1, arr)
	fmt.Printf("after changed slice %v\n", tmp1)
	// 删除第三个元素
	del_index := 3 - 1 // 默认从0开始
	tmp1 = append(arr[:del_index-1], arr[del_index:]...)
	fmt.Printf("after delete slice1 %v\n", tmp1)
}

/*

new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：
这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体 它相当于 &T{}。

make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel

*/

func new_make_diff() {
	p_1 := new([]int)
	v_1 := make([]int, 0)
	fmt.Printf("the element %v, len %d, cap %d\n", *p_1, len(*p_1), cap(*p_1))
	fmt.Printf("the element %v, len %d, cap %d\n", v_1, len(v_1), cap(v_1))
}

func main() {
	raw_arr := [10]int{1, 2, 4, 8, 7, 10, 12} // cap = 10
	//raw_slice1 := raw_arr[:] ;
	raw_slice2 := raw_arr[0 : len(raw_arr)-1]
	//fmt.Println("before update ", raw_slice1)
	fmt.Println("before update ", raw_slice2)
	//upd_slice(raw_slice1) ;
	upd_slice(raw_slice2)
	//fmt.Println("after update ", raw_slice1)
	fmt.Println("after update ", raw_slice2)
	fmt.Println("raw array ", raw_arr)
	// 测试跨界
	access_slice(raw_arr[:])
	access_slice(raw_arr[:2])
	// slice初始化操作
	init_slice()
	// slice 操作 copy/delete
	manipulate_slice(raw_arr[:])
	// make diff
	new_make_diff()
}
