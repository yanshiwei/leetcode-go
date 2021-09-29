func isPalindrome(x int) bool {
    var xStr=strconv.Itoa(x)
    if len(xStr)==1{
        return true
    }
    var i=0
    var j=len(xStr)-1
    for i<j{
        if xStr[i]!=xStr[j]{
            return false
        }
        i++
        j--
    }
    return true
}
