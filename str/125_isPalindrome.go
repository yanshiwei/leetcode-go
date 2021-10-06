
func isPalindrome(s string) bool {
    if len(s)<2{
        return true
    }
    newS:=strings.ToUpper(s)
    var buffer bytes.Buffer
    for i:=range newS{
        if unicode.IsLetter(rune(newS[i]))|| unicode.IsNumber(rune(newS[i])){
            buffer.WriteByte(newS[i])
        }
    }
    fmt.Printf(buffer.String())
    finalS:=buffer.String()
    var i=0
    var j=len(finalS)-1
    for i<j{
        if finalS[i]!=finalS[j]{
            return false
        }
        i++
        j--
    }
    return true
}
