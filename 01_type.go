package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

/*
	bool, string

	u(int), u(int8), u(int16), u(int32), u(int64), uintptr 指针

    byte, rune 字符(没有char，相当于char)

	float32, float64, complex64, complex128
*/

func main() {
	// | 3*i + 4 | = 5
	fmt.Println(cmplx.Abs(3i + 4))
	// 欧拉公式
	// e^iπ + 1 = 0
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) // Exp 是以e为底

	formate := "2006-01-02 15:04:05 Mon"
	if localLoc, err := time.LoadLocation("Local"); err != nil {
		fmt.Println("error!!")
	} else {
		fmt.Println(time.Now().In(localLoc).Format(formate))
	}
}
