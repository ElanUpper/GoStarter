package main

import (
	"fmt"
	"reflect"
)

/*
type Cat interface {
	Wow() ;
}

type Dragon struct {
	Crawl string ;
}

func (drg *Dragon) Wow() {
	fmt.Println("Hello Dragon") ;
}

func CreaterNilDrag() Cat {
	var tester Dragon ;
	return &tester ;
}

func CreaterRealDrag() Cat {
	var tester *Dragon = &Dragon{"Real Dragon"} ;
	return tester ;
}


func main ()  {

  if v := CreaterNilDrag() ; v == nil {
  	fmt.Println("nil")
  } else {
  	fmt.Println(reflect.TypeOf(v), reflect.TypeOf(v).Kind())
  }

	if v := CreaterRealDrag() ; v == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(reflect.TypeOf(v), reflect.TypeOf(v).Kind())
	}

  }
*/

type ifa interface {
	foo()
}

type test struct {
	V int
}

// 如果接收器不是指针，则ifa接口可以是指针，也可以是对象，否则只能是指针
func (t test) foo() {
	fmt.Println("fooooo:", t.V)
}

func main() {
	var aa ifa = &test{1234}
	var ap = &aa
	var o = (*ap).(*test)
	var o2 = (*ap).(*test)
	fmt.Println("main o.v:", o.V)
	o2.V = 12
	fmt.Println("main o.v:", o.V)
	fmt.Printf("main %p \n", &o.V)
	foo(ap)
	fmt.Println("main end o.v:", o.V)
	aa.foo()
}

func foo(f *ifa) {
	fmt.Printf("f: type=%v,  kind=%v \n", reflect.TypeOf(f), reflect.TypeOf(f).Kind())
	if o, ok := (*f).(*test); ok {
		fmt.Printf("%p  \n", &o.V)
		fmt.Println("foo before o.v:", o.V)
		o.V = 11
		fmt.Println("foo after  o.v:", o.V)
	}
	(*f).foo()
}