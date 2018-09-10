package main

import "fmt"

type queue []string

// push
func (que *queue) push(val string) {
	*que = append(*que, val)
	fmt.Println("push:current adress", &que)
}

// pop
func (que *queue) pop() {
	//fmt.Println(que.IsEmpty())
	if que.IsEmpty() {
		fmt.Print("queue is empty")
		return
	} else {
		fmt.Printf("%v is out of queue; ", (*que)[0])
	}
	*que = (*que)[1:]
	fmt.Println("pop:current adress", &que)
}

func (que *queue) IsEmpty() bool {
	//fmt.Println(que, *que)
	return len(*que) == 0
}

func main() {

	// 指针定义 : 常量
	var pInt *int
	vInt := 10
	pInt = &vInt
	fmt.Println(*pInt, pInt, &pInt)
	// 指针定义 : queue
	qInt := []int{10, 20}
	qdInt := &qInt
	fmt.Println(qInt, &qInt, &qdInt)

	if []string{} == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	q1 := queue{"a1"}
	q1.push("a2")
	q1.push("a3")
	q1.pop()
	q1.pop()
	q1.pop()
	q1.pop()
}
