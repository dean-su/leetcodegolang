func decodeString(s string) string {
    var countStack []int
    var strStack []string
    var k int 
    var currStr string
    
    for _, ch := range s {
        if unicode.IsDigit(ch) {
            k = k * 10 + int(ch - '0'); 
        } else if ch == '[' {
            countStack = append(countStack, k)
            strStack = append(strStack, currStr)
            k = 0 
            currStr = ""
        } else if ch == ']' {
            n, m := len(strStack)-1, len(countStack)-1
            temp := strStack[n]
            strStack = strStack[:n]
            for i := countStack[m]; i >0 ;i-- {
                temp = temp + currStr
            }
            countStack = countStack[:m]
            currStr = temp
        } else {
            currStr = currStr + string(ch) 
        }
      
    }
    return currStr
    
}