package main

import "fmt"

// queue是一个可以承载所有类型的切片
type queue []interface{}

// push
func (que *queue) push(val interface{}) {
	switch val.(type) { // 仅仅支持int&string
	case int:
		*que = append(*que, val)
	case string:
		*que = append(*que, val)
	default:
		fmt.Printf("not support!\n")
		return
	}
	fmt.Println("push:current adress", &que)
}

// pop
func (que *queue) pop() queue {
	if que.IsEmpty() {
		fmt.Printf("queue is empty\n")
		return nil
	} else {
		// 打印队列中不同类型对应的值
		switch (*que)[0].(type) {
		case int:
			fmt.Printf("int: %v is out of queue.\n", (*que)[0].(int))
		case string:
			fmt.Printf("string: %v is out of queue.\n", (*que)[0].(string))
		default:
			fmt.Printf("unknow: %v is out of queue.\n", (*que)[0])
		}
		return (*que)[1:]
	}
}

func (que *queue) IsEmpty() bool {
	//fmt.Println(que, *que)
	return len(*que) == 0
}

func main() {

	q1 := queue{11.222} // 初始化填入float类型
	q1.push(10)
	q1.push("a2")
	q1.push("a3")
	q1 = q1.pop()
	q1 = q1.pop()
	q1 = q1.pop()
	q1 = q1.pop()
}
