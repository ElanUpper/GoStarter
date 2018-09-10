package test

import (
	. "07_test"
	"testing"
)

func testCase(t *testing.T) {

	testData := []struct{ a, b, c int }{
		{10, 10, 20},
		{15, 20, 35},
		{20, 25, 45},
		{25, 30, 55},
	}

	for _, val := range testData {
		if Act := Calc_add(val.a, val.b); Act != val.c {
			t.Errorf("Calc_add(%v, %v) != the expected value %v ",
				val.a, val.b, val.c)
		} else {
			t.Logf("Calc_add(%v, %v) == the expected value %v ",
				val.a, val.b, val.c)
		}
	}
}

func BenchmarkCase(t *testing.B) {

	testData := []struct{ a, b, c int }{
		{10, 10, 20},
		{15, 20, 35},
		{20, 25, 45},
		{25, 30, 55},
	}

	// 采用系统合适的计算次数
	for i := 0; i < t.N; i++ {
		for _, val := range testData {
			Calc_add(val.a, val.b)
		}
	}
}
