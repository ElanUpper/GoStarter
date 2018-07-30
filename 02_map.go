package main

import "fmt"

func map_int(){
	// var进行初始化  key为字符串 value为整数型
	var map_1 = map [string]int{}
	map_2 := make(map [string]int)
	//
	map_3 := map [string]int {
		"s1": 10,
		"s2": 20,
		"s3": 30,
	}
	fmt.Println(map_1, map_2, map_3)
	// 打印各个元素的列表
	for k, v := range(map_3) {
		fmt.Printf("key %s value %d\n", k, v)
	}
}

func get_longest_str(str string)(int, int, string){
	// 定义一个key为string value为string的字典
	strPos := 0
	maxLen := 0
	strMap := map [string]int {} ;
	for index, value := range (str){
		// 如果出现重复字母的话 那么就用这个字母后面的位置为起点  (这里添加增强判断是否字段key存在)
		if mapVal, ok_code := strMap[string(value)]; ok_code && strPos <= mapVal + 1 {
			strPos = mapVal + 1
		}
		// 如果存储长度大于最大长度 那么使用当前存储长度
		if index - maxLen + 1 > maxLen {
			maxLen = index - maxLen + 1
		}
		// 每次存储字符的最新位置
		strMap[string(value)] = index
	}
	fmt.Println(strMap,  str[strPos:maxLen])
	// 返回起始位置  长度  无重复的字符串
	return strPos, maxLen, str[strPos:strPos+maxLen];
}


func main(){
	map_int();
	fmt.Println(get_longest_str("helloworld"))
	fmt.Println(get_longest_str("我爱go语言"))
}