func max(a, b int) int {
    if a <= b {
        return b
    }
    return a
}

func lengthOfLongestSubstring(s string) int {
    	ret := 0
	mc := make(map[rune]int)
	i := 0
	for j, v := range s {
		if _, f := mc[v]; f {
			i = max(mc[v], i)
			mc[v] = j+1
		}
		
		ret = max(ret, j-i+1)
		mc[v] = j+1
	}
	return ret
}