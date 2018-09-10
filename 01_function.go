package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func multi_returned_value(num_a, num_b int) (int, int, string) {
	return num_a % num_b, num_a / num_b, "no error"
}

func multi_returned_value_err(num_a, num_b int, opertor string) (int, error) {
	switch opertor {
	case "+":
		return num_a + num_b, nil
	case "-":
		return num_a - num_b, nil
	case "*":
		return num_a * num_b, nil
	case "/":
		if num_a != 0 {
			return num_a / num_b, nil
		} else {
			return 0, fmt.Errorf("dividend equal zero!")
		}
	default:
		return 0, fmt.Errorf("unsupported operator!")
	}
}

func apply_func_as_parameter(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()    // 获取函数指针
	f_name := runtime.FuncForPC(p).Name() // 获取函数名称
	fmt.Printf("the funcition %s with parameter(%d, %d) is running!\n", f_name, a, b)
	return op(a, b)
}

func f_power(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func f_add(a int, b int) int {
	return a + b
}

// 不定参数传递
// 返回参数列表数字和
func f_variable_params(members ...interface{}) int {
	s := 0
	for i, v := range members {
		switch t := reflect.TypeOf(v); t.Kind() {
		case reflect.Int:
			fmt.Printf("step %d, value %d\n", i, members[i].(int))
			s += members[i].(int)
		case reflect.Float64:
			//fmt.Printf("step %d, value %f\n", i, v.(float64))
			s += int(v.(float64))
		default:
			panic("not support !")
		}

	}
	return s
}

func main() {

	// 使用匿名函数+调用参数
	fmt.Println(
		func(args ...int) int {
			SumVal := 0
			for i, v := range args {
				fmt.Printf("%d, %d, %d\n", i, v, args[i])
				SumVal += v
			}
			return SumVal
		}(10, 20, 30, 40))

	// _ 表示这个返回值 我们这里不需要
	a, b, _ := multi_returned_value(10, 3)
	fmt.Printf("the value %d, %d\n", a, b)
	if result, err := multi_returned_value_err(1, 10, "0"); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("the result %d\n", result)
	}
	// 这里我们使用函数式编程
	fmt.Printf("the result is %d\n", f_add(f_power(10, 2), 10))
	// 函数映射 动态使用函数
	//func_opt := "f_power"
	fmt.Printf("the result is %d\n", apply_func_as_parameter(f_power, 10, 2))
	// 使用匿名函数
	fmt.Printf("the result is %d\n", apply_func_as_parameter(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 10, 2))

	// 使用可变参数
	fmt.Printf("variable params: the result is %d\n", f_variable_params(1, 2, 3, 4, 5))
	fmt.Printf("variable params: the result is %d\n", f_variable_params(1.2, 2.4, 3.6))
}
