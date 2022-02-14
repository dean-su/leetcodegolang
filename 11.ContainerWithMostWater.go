func min(a,b int) int{
    if a < b {
        return a
    }else{
        return b
    }
}
func max(a,b int) int{
    if a < b{
        return b
    }else{
        return a
    }
}

func maxArea(height []int) int {
    l, r := 0, len(height)-1
    res := 0
    for l < r {
        h := min(height[l], height[r])
        res = max(res, h*(r-l))
        if(height[l]<height[r]) {
            l++
        }else {
            r--
        }
    }
    return res
}