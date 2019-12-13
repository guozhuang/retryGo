package main

import "fmt"

func getNoRepeatStrLen(s string) int {
	maxLen := 0

	lastCurrent := make(map[rune]int)
	start := 0

	//先针对byte，后面针对rune，对应修改lastcurrent的key类型即可
	for i, ch := range []rune(s) {
		lastI, ok := lastCurrent[ch]

		//这里表示在之前一段进行匹配
		if ok && lastI >= start {
			//说明重新开始标记长度
			start = lastI + 1
		}

		//查看长度更新
		if maxLen < (i - start + 1) {
			maxLen = i - start + 1
		}

		lastCurrent[ch] = i //进行迭代
	}

	return maxLen
}

func main() {
	testStr := "abcabc"
	fmt.Println("abcabc :", getNoRepeatStrLen(testStr))

	testStr1 := "bbb"
	fmt.Println("bbb :", getNoRepeatStrLen(testStr1))

	testStr2 := "黑暗中有什么不是黑暗的呢！"
	fmt.Println("", getNoRepeatStrLen(testStr2))
}
