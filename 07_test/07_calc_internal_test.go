package calcer

import "testing"

var testData = []struct{ a, b, c int }{
	{10, 10, 20},
	{15, 20, 35},
	{20, 25, 45},
	{25, 30, 55},
}

func testCase(t *testing.T) {

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

	// 采用系统合适的计算次数
	for i := 0; i < t.N; i++ {
		for _, val := range testData {
			Calc_add(val.a, val.b)
		}
	}
}

// 使用工具提速
// go test -bench . -cpuprofile cpu.out
// 安装Graphviz  设置好path = Graphviz + Graphviz/bin
// go tool pprof cpu.out
//  (pprof) web -> 打开运行时图
