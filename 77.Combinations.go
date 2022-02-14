func loop(i, n, k int, temp []int, res *[][]int){
    if len(temp) == k {
        cp := make([]int, len(temp))
        copy(cp, temp)
        *res = append(*res, cp)
        return
    }
    if i>n {
        return
    }
    temp = append(temp, i)
    loop(i+1, n, k, temp, res)
    temp = temp[:len(temp)-1]
     loop(i+1, n, k, temp, res)
}

func combine(n int, k int) [][]int {
    res := [][]int{}
    temp := make([]int, 0 , k)
    loop(1, n, k, temp, &res)
    return res
}