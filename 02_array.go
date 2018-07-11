package main

import "fmt"

func arr_init(){
	var arr1 [3]int ;
	arr2 := [3]int{1, 2, 3} ;
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2, arr3)
	for i, v := range arr3 {
		fmt.Println(i, v, arr3[i])
	}
}

// print 5 elements
func arr_print_elements(arr [5]int) {
	for k, v := range( arr ) {
		fmt.Printf("  %d:%d\n", k, v)
		if v > 5 {
			arr[k] = -100 ;  // 因为go语言中是值传递 所以不会修改数组
		}
	}
}

// 指针交换值
func arr_pointer_exchange_max(arr *[5]int){
	max_i, max_v := 0, 0
	for k, v := range( arr ) {
		if max_v < v {
			max_i, max_v = k, v
			if max_v > 5 {
				arr[k] = 100 ;  // 如果大于5的话 就修改数组中的值  如果需要修改那么需要使用指针
			}
		}
	}
	fmt.Printf("max index: %d, max value %d\n", max_i, max_v)
}

func main(){
	// arr_init()
	arr1 := [...]int{1, 2, 3, 4, 7} ;
	arr_print_elements(arr1)  // 值传递
	arr_pointer_exchange_max(&arr1) // 指针传递
	fmt.Println(arr1)
}
