func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    i1, i2 := 0,0 // left pointers
    n, m := len(nums1), len(nums2)
    j1, j2 := n-1, m-1 // right pointers
    total := n + m // total size
    for ;total > 2; total -= 2{
        // move from left
        if i1 <= j1 && i2 <= j2 {
            if nums1[i1] <= nums2[i2] {
                i1++
            } else {
                i2++
            }
        } else if i1 <= j1{
            i1++
        } else if i2 <= j2{
            i2++
        }
        
        // move from right
        if j1 >= i1 && j2 >= i2 {
            if nums1[j1] <= nums2[j2] {
                j2--
            } else {
                j1--
            }
        } else if j1 >= i1{
            j1--
        } else if j2 >= i2{
            j2--
        }
    }

    val, cnt := 0, 0
    // add remaining items (1 or 2) and keep sum and count
    for c := i1;c <= j1; c++ {
        val += nums1[c] 
        cnt++
    }
    for c := i2;c <= j2; c++ {
        val += nums2[c] 
        cnt++
    }
    return float64(val)/float64(cnt)
}