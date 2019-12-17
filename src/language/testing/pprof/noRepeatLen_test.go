package main

import "testing"

//先写一个单元测试
func TestGetNoRepeatStrLen(t *testing.T) {
	tests := []struct {
		str string
		len int
	}{
		{"ababc", 3},
		{"hello", 3},
		{"bbb", 1},
		{"黑暗中有什么不是黑暗的呢!", 11},
	}

	for _, tv := range tests {
		testR := getNoRepeatStrLen(tv.str)

		if testR != tv.len {
			t.Errorf("getNoRepeatStrLen(%s), got %d, expect %d", tv.str, testR, tv.len)
		}
	}
}

//写一个benchmark
func BenchmarkGetNoRepeatStrLen(b *testing.B) {
	testStr := "黑暗中有什么不是黑暗的呢!"
	expLen := 11

	for i := 0; i < b.N; i++ {
		testR := getNoRepeatStrLen(testStr)

		if testR != expLen {
			b.Errorf("getNoRepeatStrLen(%s), got %d, expect %d", testStr, testR, expLen)
		}
	}
}

//【使用ide进行压测或者pprof都可以进行处理】
//命令行压测：go test -bench .

//进行pprof检测:go test -bench . -cpuprofile cpu.out
//生成的cpu.out是一个二进制文件，使用: go tool pprof cpu.out来解析【使用web进行打开，需要安装python的画图库】

//解析结果发现map和decoderune操作。【要明白耗时点是否必要处理】
//如果觉得map性能占比不好，将map置为数组，复杂度肯定会降低【但是会触发垃圾回收】
//map适合对处理字符串长度不长的情况下，进行处理，而长度大的时候，使用slice来进行处理
