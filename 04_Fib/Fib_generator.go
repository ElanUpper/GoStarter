package _4_Fib

func Fib_gen() func() int {
	num_f, num_a := 0, 1
	return func() int {
		num_f, num_a = num_a, num_a+num_f
		return num_f
	}
}

func FunClosure() func() int {
	num_a, num_b := 1, 2
	return func() int {
		num_a, num_b = num_b, num_a+num_b
		return num_b
	}
}
