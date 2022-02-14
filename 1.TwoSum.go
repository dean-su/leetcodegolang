func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        if v, f := numMap[target - num]; f {
            return []int{v, i}
        }
        numMap[num] = i
    }
    return []int{}
}