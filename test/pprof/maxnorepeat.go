package pprof

func NoRepeat(s string) int {
	lastCur := make(map[rune]int) //标记旧有字符最后出现的位置
	start := 0
	maxLen := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastCur[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		//最大值更新
		if i-start+1 > maxLen {
			maxLen = i - start + i
		}

		lastCur[ch] = i
	}

	return maxLen
}
