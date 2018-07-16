package main

import (
	"fmt"

	"sort"
	"unicode/utf8"
)

// 除了slice map function以外，内建类型都可以作为key
func map_get_value(){
	dict_1 := map[string]string{
		"hello1" : "print1",
		"hello2" : "print2",
		"hello3" : "print3",
	}
	// get value by key， 如果okcode=True，那么就代表这个key存在，否则代表不存在，返回zero value
	if val, okcode := dict_1["hello1"]; okcode{
		fmt.Printf("the value is %s\n", val)
	} else {
		fmt.Println("the key doesn't exist")
	}
	// delete key
	fmt.Println("after delete the key: hello1")
	delete(dict_1, "hello1")
	if val, err := dict_1["hello1"]; err{
		fmt.Printf("the value is %s\n", val)
	} else {
		fmt.Println("the key doesn't exist")
	}
}

func map_get_sorted_value(){
	dict_0 := map[string]string{
		"10" : "print1",
		"13" : "print2",
		"09" : "print3",
		"07" : "print3",
		"16" : "print3",
	}
	// 字典转换为list
	/*
	vals  := reflect.ValueOf(dict_0)
	keys  := vals.MapKeys()
	*/
	keys := []string{} // 创建string list
	for item := range dict_0 {
		keys = append(keys, item)
	}
	sort.Strings(keys)
	for _, v := range(keys) {
		fmt.Printf("key %s, value %s \n", v, dict_0[v])
	}

	/*
	list1 := []string{}
	for _, item := range keys {
		// 两种方法获取类型
		//fmt.Printf(":%T\n", item)
		//fmt.Println(reflect.TypeOf(item))
		// 将reflect.Value 转换为string
		list1 = append(list1, reflect.ValueOf(item).Interface().(string) )
	}
	*/

}

func get_max_chars(str string) int {
	// 打印字符串每一个字母
	/*list_count, list_save := 0, []string {}
	for _, v := range str {

		list_save = append(list_save, string(v)) ;

		fmt.Printf("%s", string(v))
	}
	*/
	lastOccurred := make(map[string]int)
	start 	  := 0
	maxLength := 0
	lv_str 	  := ""

	for i, ch := range []byte(str){
		lv_str = string(ch)
		fmt.Println(i, ch, lv_str)
		if lastOccurred[lv_str] >= start {

			start = lastOccurred[lv_str] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[lv_str] = i ;
	}
	fmt.Println(lastOccurred)
	return maxLength ;
}

// 解析中文
func decode_chinese_words(str string){
	// 16进制显示 每一个字符 还原unicode字符集
	for k, val := range str{
		fmt.Printf("index %d, value %X\n", k, val)
	}
	// 转换为rune数组
	for rune_index, rune_item := range []rune(str) {
		fmt.Printf("index %d, value %c\n", rune_index, rune_item) ;
	}
	// 转换为byte数组
	str_bytes := []byte(str)
	for len(str_bytes) > 0 {
		// 解析每个字符  内容&长度
		ch, size := utf8.DecodeRune(str_bytes)
		// 剪切字符
		str_bytes = str_bytes[size:]
		fmt.Printf("size:%d, char:%c\n", size, ch)
	}

}

func find_longest_str(str string) (int, int, string, map[rune]int) {
	map_str := map [rune]int{} ;
	start_pos  := 0
	start_out  := 0
	max_Length := 0
	for i, ch := range []rune(str){
		// 如果发生重复
		if val, ok_code := map_str[ch]; ok_code && start_pos <= val + 1 {
			start_pos = val + 1
		}
		if max_Length < i - start_pos + 1 {
			start_out  = start_pos  // 仅仅找到更长字符串的时候 修改字符串起始位置
			max_Length = i - start_pos + 1
		}
		map_str[ch] = i // 存储每次ch对应i的值
		fmt.Printf("%d, %c\n", i, ch)
	}
	fmt.Printf("start %d, max length: %d, string %c, %v\n", start_out, max_Length,
											[]rune(str)[start_out:max_Length+start_out], map_str )
	  							  // [] 取值范围是[)
	return start_pos, max_Length, str[start_pos:max_Length+start_pos-1], map_str ;
	// "start %d, max length: %d, string %s, %v"
}

func main(){
	//map_get_value()
	//map_get_sorted_value()
	//fmt.Println(get_max_chars("12301234043216"))
	//decode_chinese_words("hello,王阳!!")
	find_longest_str("我爱你你爱我吗我的")
}