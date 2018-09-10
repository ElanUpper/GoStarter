package main

import "fmt"

func have_expression_after_switch(score int) string {
	var grade_level = ""
	switch score {
	// 默认是自动break的，如果不加语句，那么就是什么都不做，而不是按照下面语句进行执行！！！
	// 如果需要那么需要添加fallthrough (类似于3)
	case 0:
	case 1:
		grade_level = "C"
	case 2:
		grade_level = "C+"
	case 3:
		fallthrough // fallthrough: don't breadk
	case 4, 5, 6:
		grade_level = "B"
	// 即使赋值也会被下面语句覆盖 而不会直接返回
	case 7, 8, 9:
		grade_level = "B+"
		fallthrough
	case 10:
		grade_level = "A"
	default:
		panic("error")
	}

	return fmt.Sprintf("score: %d, level: %s\n", score, grade_level)
}

// switch多条件使用
func no_expression_after_switch(score int, auth int) string {
	var grade_level = ""
	switch {
	case score < 0 || score > 100:
		// panic(
		fmt.Sprintf("illege number is %d", score)
		// )
	case score <= 60:
		grade_level = "D"
	case score <= 80:
		grade_level = "C"
	case score <= 90:
		grade_level = "B"
	case score < 100:
		grade_level = "A"
	case score <= 100 && auth == 2:
		grade_level = "A+"

	}

	return grade_level
}

func if_expression(price int, quatity int) string {
	if gross := price * quatity; gross > 1000 {
		return "qualified"
	} else {
		return "below the mark"
	}
}

func main() {
	fmt.Println(if_expression(100, 20))
	fmt.Println(
		have_expression_after_switch(0),
		have_expression_after_switch(3),
		have_expression_after_switch(7),
		have_expression_after_switch(10))
	fmt.Println(no_expression_after_switch(100, 2))
}
