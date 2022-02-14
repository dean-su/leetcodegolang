func searchRange(nums []int, target int) []int {
    return []int{firstPos(nums, target), lastPos(nums, target)}
}

func firstPos(nums []int, target int) int {
    l, r, m := 0, len(nums), 0
    for l < r {
        m = (l+ r)/2
        if nums[m] >= target {
            r = m
        } else {
          l = m+1  
        }
    }
    
    if l == len(nums) || nums[l] != target {
        return -1
    }
    
    return l
}

func lastPos(nums []int, target int) int {
    l, r, m := 0, len(nums), 0
    for l < r {
        m = (l+r)/2
        if nums[m] > target {
            r = m
        } else {
          l = m+1  
        }
    }
    
    l--
    if l < 0 || nums[l] != target {
        return -1
    }
    
    return l
}