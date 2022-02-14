func max(a,b int) int {
    if a < b {
        return b
    }else {
        return a
    }
}
func maxSubArray(nums []int) int {
    //nums = [-2,1,-3,4,-1,2,1,-5,4]
    //f =    [-2,1,-2,4,3,5,6,1,5]
    //f[i] = f[i-1] > 0? f[i-1] + nums[i] : nums[i]
    f := make([]int, len(nums))
    f[0] = nums[0]
    m := nums[0]
    for i := 1; i < len(nums); i++{
        f[i] = max(f[i-1] + nums[i], nums[i])
        
        if m < f[i]{
            m = f[i]
        }
    }
    
    
 
    return m
}