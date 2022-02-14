func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    
    for l < r {
        m := (l+r)/2
        if nums[m] == target {return m}
        
        if nums[m] >= nums[l] {
            if target < nums[m] && target >= nums[l] {
                r = m
            }else{
                l = m+1
            }
            
        }else {
            if target > nums[m] && target <= nums[r] {
                l = m+1
            }else{
                r = m
            }
        }
    }
    
    if nums[l] == target {
        return l
    }else if nums[r] == target {
        return r
    }else {
        return -1
    }
}