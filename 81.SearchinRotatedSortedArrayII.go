/*
func search(nums []int, target int) bool {
    if len(nums) == 1{
        // base case
        return nums[0] == target
    }
    if search(nums[:len(nums)/2],target) || search(nums[len(nums)/2:],target){

        return true
    }
    return false
}
*/
func search(nums []int, target int) bool {
    l , r := 0, len(nums)-1
    
    for l+1<r {
        m := (l+r)/2
        if nums[m] == target {return true}
        
        if nums[l] == nums[m] && nums[m] == nums[r] {
            l++
            r--
            continue
        }
        if nums[m] >= nums[l] {
            if target < nums[m] && target >= nums[l] {
                r = m
            }else {
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
        
    
    if nums[l] == target || nums[r] ==target {
        return true
    }else {
        return false
    }
}