package main

import "fmt"

func have_expression_after_switch(score int) string {
	var grade_level = "";
	switch score {
	case 0:
	case 1:
	case 2:
		grade_level = "C"
		fallthrough
	case 3:
		grade_level = "B"
		fallthrough
	case 4:
	case 5:
		grade_level = "A"
	default:
		panic("error")
	}

	return grade_level;
}

// switch多条件使用
func no_expression_after_switch(score int, auth int) string {
	var grade_level = "";
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

	return grade_level;
}

func if_expression(price int, quatity int) string {
	if gross := price * quatity ; gross > 1000 {
		return "qualified"
	} else {
		return "below the mark"
	}
}

func main() {
	fmt.Println(if_expression(100, 20))
	fmt.Println(have_expression_after_switch(4),
		have_expression_after_switch(3),
		have_expression_after_switch(1))
	fmt.Println(no_expression_after_switch(100, 2))
}
