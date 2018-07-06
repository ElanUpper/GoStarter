package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

/*
	bool, string

	u(int), u(int8), u(int16), u(int32), u(int64), uintptr 指针

    byte, rune 字符(char)

	float32, float64, complex64, complex128
 */

func main(){
	// | 3*i + 4 | = 5
	fmt.Println( cmplx.Abs(3i + 4) )
	// 欧拉公式
	// e^iπ + 1 = 0
	fmt.Println( cmplx.Pow( math.E, 1i * math.Pi ) + 1 )
	fmt.Printf( "%.3f", cmplx.Exp( 1i * math.Pi ) + 1 )  // Exp 是以e为底
}
