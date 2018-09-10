package main

import (
	"fmt"
	"reflect"
	"time"
)

/*

// dock typing

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

func chann_test() {
	c := make(chan int, 1)
	c <- 10
	close(c)
	v, ok := <-c
	if ok {
		fmt.Println(v, ok)
		v, ok = <-c
		fmt.Println(v, ok)
		c <- 20
	} else {
		println("closed!")
	}
}

func fib_test() func() int {
	n_a, n_b := 0, 1
	return func() int {
		n_a, n_b = n_b, n_a+n_b
		return n_a
	}
}

func main() {

	formate := "2006-01-02 15:04:05 Mon"
	if localLoc, err := time.LoadLocation("Local"); err != nil {
		fmt.Println("error!!")
	} else {
		fmt.Println(time.Now().In(localLoc).Format(formate))
	}

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

	inst_fib := fib_test()
	for i := 0; i < 10; i++ {
		fmt.Println(inst_fib())
	}
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
