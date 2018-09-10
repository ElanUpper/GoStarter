package main

import "fmt"

//*---------------------- 接口定义
type inter1 interface {
	SetData(data map[string]string)
}

type inter2 interface {
	GetData() string
}

type inter interface {
	inter1
	inter2
	//auth(id int); // 打印验证后用户
}

// test function search the implement
type simpleMap struct {
	mapper map[string]string
}

func (Map *simpleMap) SetData(data map[string]string) {
	Map.mapper = data
}

func (Map *simpleMap) GetData() string {
	return "hello world"
}

//*---------------------- 接口实现
type DataCenter struct {
	userId int
	Record string
}

// 重新定义string方法
func (inter *DataCenter) String() string {
	return fmt.Sprintf("User id: %d, Record %s\n", inter.userId, inter.Record)
}

func (inter *DataCenter) SetData(str map[string]string) {
	if content, ok := str["content"]; ok {
		inter.Record = content
	}
}

func (inter *DataCenter) GetData() string {
	return inter.Record
}

func (inter *DataCenter) auth(id int) {
	inter.userId = id
}

//*---------------------- 接口调用
func reader_center(contr inter) string {
	return contr.GetData()
}

func writer_center(contr inter, rec map[string]string) {
	contr.SetData(rec)
}

func auth_center(contr inter) {

}

func main() {
	var data1 inter
	data1 = &DataCenter{10, ""}
	writer_center(data1, map[string]string{
		"name":    "elan",
		"content": "only name",
	})
	fmt.Println(reader_center(data1))
	fmt.Print(data1)
}
