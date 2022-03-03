/*
func mySqrt(x int) int {
    if x == 0 || x == 1{
        return x
    }
    var l, r int64
    l = 0
    r = int64(x)
    t := int64(x)
    for l < r {
        m := (l+r)/2
        if m*m > t {
            r = m
        }else {
            l = m +1
        }
    }
    return int(l-1)
    
}
*/

func mySqrt(x int) int {
    r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
