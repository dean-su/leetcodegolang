/*
func searchInsert(nums []int, target int) int {
    for i := 0; i<len(nums); i++ {
        if nums[i] >= target{
            return i
        }
    }
    return len(nums)
}
=====================
func searchInsert(nums []int, target int) int {
    l , r := 0, len(nums)
    var m int
    for l<r {
        m = l + (r-l)/2
        if nums[m] == target {
            return m
        }else if nums[m]>target {
            r = m
        } else {
            l = m + 1
        }
    }
    
    return l
}
*/
func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int)bool{
		 return nums[i] >= target
	 })
 
 }