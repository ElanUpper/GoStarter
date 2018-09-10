package main

import (
	"errors"
	"fmt"
	"log"
)

func child() {

	// catch panic
	defer func() {
		fmt.Println("child: starting...")
		if handle_err := recover(); handle_err != nil {
			if err, ok := handle_err.(error); ok {
				fmt.Printf("child: the error info %v \n", err.Error())
			} else {
				// unknow type/empty
				errInfo := fmt.Sprintf("Child: I don't konw how to handle it %v \n", handle_err)
				log.Printf(errInfo)
				panic(handle_err)
			}
		}
		fmt.Println("child: end !!")
	}()

	func() {

		fmt.Println("subroutine starting ..... ")

		// 协程抛出异常  外部无法recover 仅仅只能在内部处理
		go func() {
			defer func() {
				// sub routine: catch
				handle_err := recover()
				log.Printf("subroutine: %v\n", handle_err)
			}()
			//panic("hello world")
			panic(errors.New("error: from subroutine"))
		}()

		// 子进程抛出异常
		//zero := 0
		////panic( 1 / zero )
		//panic(zero)

		fmt.Println("subroutine end !!")

	}()

}

func main() {

	// catch panic
	defer func() {
		if handle_err := recover(); handle_err != nil {
			if err, ok := handle_err.(error); ok {
				fmt.Printf("main: the error info %v\n", err.Error())
			} else {
				log.Printf("main: I don't have ability to handle it %v\n", handle_err)
			}
		}
	}()

	child()

	fmt.Println("main: done!")

}
