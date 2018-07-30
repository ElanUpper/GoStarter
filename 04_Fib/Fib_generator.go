package _4_Fib

func Fib_gen() func() int {
	num_f, num_a := 0, 1
	return func() int {
		num_f, num_a = num_a, num_a + num_f ;
		return num_f
	}
}

