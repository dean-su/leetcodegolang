type NumArray struct {
    sums_ []int
}


func Constructor(nums []int) NumArray {
    n := len(nums)
    if n == 0 {
        return NumArray{[]int{}}
    }
    sums := make([]int, n)
    sums[0] = nums[0]
    for i :=1; i < n; i++ {
        sums[i] = sums[i-1] + nums[i]
    }
    
    return NumArray{sums}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 { return this.sums_[right] }
    return this.sums_[right] - this.sums_[left-1]
}