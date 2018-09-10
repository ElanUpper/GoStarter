package main

import (
	"fmt"
	"reflect"
)

type numb int ;

// 创建一个有自增方法的int类型
func (num *numb) addNumb(iAdd int) int {
	*num += numb(iAdd) ;
	fmt.Println("the address:", num)
	retVal := int(*num)
	return retVal
}

// type switch
func GetType (inter interface{}) {
	switch tp := inter.(type) {
	case numb:
		fmt.Println("numb", tp.addNumb(100) ) ;
	case int:
		fmt.Println("numb", tp ) ;
	case string :
		fmt.Println( tp ) ;
	default:
		fmt.Println("unsupported", reflect.TypeOf(inter))
	}
}

// reflace version
func GetType_reflect (inter interface{}) {
	switch tp := reflect.TypeOf(inter); tp.Kind() {
	case reflect.Struct :
		fmt.Println(tp.Name())
	case reflect.Int : // int:
		switch tp.Name() {
			case "numb" : // 虽然我们的类型numb实现了一个方法，但是其本质还是int
				Snumb := inter.(numb) ;
				Pnumb := &Snumb
				fmt.Println("numb", Pnumb.addNumb(100) ) ;
			default:
				fmt.Println("int", inter.(int) ) ;
		}
	case reflect.String :
		fmt.Println("numb", inter.(string) ) ;
	default:
		fmt.Println("unsupported", reflect.TypeOf(inter), tp.Kind())
	}
}

// reflect解析结构
type emp struct {
	Id    int ;
	Name  string;
}

func (e emp)GetName() string {
	return e.Name
}

func (e *emp)SetName(name string)  {  // 为什么会被忽略掉！！！！
	e.Name = name
}

func reflect_stru_simple(e emp){  // 解析简单结构

	// type
	t := reflect.TypeOf(e) ;
	fmt.Println("structure name: ", t.Name())
	fmt.Println("type count:", t.NumField())

	// value
	v := reflect.ValueOf(e)

	// 解析基础结构
	for i:=0; i<t.NumField(); i++ {
		nam := t.Field(i).Name
		typ := t.Field(i).Type
		val := v.Field(i).Interface()
		fmt.Printf("name: %4s, type: %4s, val: %4v\n", nam, typ, val)
	}

	// 解析方法
	fmt.Println("method count:", t.NumMethod())
	for i:=0; i<t.NumMethod();i++{
		methName := t.Method(i).Name
		methType := t.Method(i).Type
		fmt.Printf("name: %4s, type: %4s\n", methName, methType)
	}
}

type company struct {
	eps 	emp ;
	Name	string ;
	Loc struct {
		left  string ;
		right string ;
	}
}

func reflect_stru_complex(c company){  // 解析简单结构

	t := reflect.TypeOf(c)

	v := reflect.ValueOf(c)

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("fields: %#v, value: %#v\n", t.Field(i), v.Field(i))
	}

	// 打印emp结构
	fmt.Println("emp: ", t.FieldByIndex([]int{0}) )
	fmt.Println("emp: ", t.FieldByIndex([]int{0, 0}) )  // 获取emp-id
	// 打印company结构 (非匿名)
	fmt.Println("Loc: ", t.FieldByIndex([]int{2}) )
	fmt.Println("Loc: ", t.FieldByIndex([]int{2, 1}) )  // 获取Loc-right
}


// -------------------------------  修改结构

type Phone struct {
	Id  	int ;
	Name 	string ;
}

func (ph *Phone) SetId(id int) {
	ph.Id = id ;
}

func (ph Phone) GetName() string {
	return ph.Name
}

// reflect 动态调用函数
func SetId_reflect(phone *Phone, id int){
	v := reflect.ValueOf(phone)
	if meth := v.MethodByName("SetId"); meth.IsValid() {
		// 设置reflect 变量
		ref_id := []reflect.Value{reflect.ValueOf(id) };
		// 调用函数
		meth.Call(ref_id) ;
	} else {
		fmt.Println("it don't contain this method")
	}
}

func SetName(phone interface{}, name string){
	v := reflect.ValueOf(phone);
	// 首先需要确定 类型属于指针 以及 值可以被修改
	if elem := v.Elem() ;v.Kind() == reflect.Ptr && elem.CanSet(){ // Elem returns the value that the interface
		// 获取Name属性，并且判断有效
		if nm := elem.FieldByName("Name"); nm.IsValid(){
			if nm.Kind() == reflect.String {
				nm.SetString(name)
			}
		} else {
			fmt.Println("it don't contain field Name !!", v.CanSet())
			return
		}
	} else {
		fmt.Println("only support pointer!!", v.Kind(), v.Elem().CanSet())
		return
	}
}

func main(){
	/*
	fmt.Println("-------------------- 类型解析 --------------------------")
	iNum := numb(0) ;
	iNum.addNumb(100)
	iNum.addNumb(100)
	iNum.addNumb(100)
	fmt.Println(iNum)

	GetType(iNum)
	GetType_reflect(10.222)

	GetType_reflect(iNum)
	GetType_reflect(10)
	*/

	fmt.Println("---------------------- 反射修改结构 ---------------------------")
	ph := Phone{10, "moto"}
	fmt.Println(ph)
	SetName(&ph, "lolo")
	fmt.Println(ph)
	SetId_reflect(&ph, 001)
	fmt.Println(ph)

	/*
	fmt.Println("---------------------- 反射解析结构 ---------------------------")
	ep := emp{Id:001, Name:"elan"}
	reflect_stru_simple(ep)

	cmp := company{ eps:emp{Id:002, Name:"elan"},
					Loc: struct {   // 结构体内 匿名结构赋值(因为外部结构无法知道type，所以外部必须也使用匿名)
						left  string ;
						right string ;
					}{"100", "150"},
					Name:"Global01"}
	reflect_stru_complex(cmp)
	*/

}

