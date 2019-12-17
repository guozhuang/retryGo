package main

import (
	"testing"
)

//简单的手动单元测试示例【表格驱动测试】【正常环境还是需要使用测试框架】
//虽然单独的add处理函数是不需要进行单元测试
func TestAdd(t *testing.T) {
	//构建测试数据集
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{5, 6, 11},
		//{math.MaxInt32, 1, math.MinInt32},//这个样例没有通过
	}

	//构建测试逻辑
	for _, tt := range tests {
		if re := add(tt.a, tt.b); re != tt.c {
			t.Errorf("func add is testing wrong:add(%d, %d), expect:%d, got %d", tt.a, tt.b, tt.c, re) //演示错误信息的标准化【将函数以及参数通过反射进行标准化？】
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	add1 := 5
	add2 := 6
	re := 11

	for i := 0; i < b.N; i++ {
		c := add(add1, add2)
		if c != re {
			b.Errorf("func add is testing wrong:add(%d, %d), expect:%d, got %d", add1, add2, re, c)
		}
	}
}
