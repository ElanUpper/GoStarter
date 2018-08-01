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

/* reflace version */
func GetType_reflect (inter interface{}) {
	switch tp := reflect.TypeOf(inter); tp {
	case numb :
		fmt.Println("numb", inter.(*numb).addNumb(100) ) ;
	case int : // int:
		fmt.Println("numb", inter.(int) ) ;
	case string :
		fmt.Println("numb", inter.(string) ) ;
	default:
		fmt.Println("unsupported", reflect.TypeOf(inter), tp.Kind())
	}
}

func main(){
	iNum := numb(0) ;
	iNum.addNumb(100)
	iNum.addNumb(100)
	iNum.addNumb(100)
	fmt.Println(iNum)

	/*
	GetType(iNum)
	GetType(10)
	GetType("hello world")
	GetType(10.222)
	*/

}

