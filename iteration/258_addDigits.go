func addDigits(num int) int {
    if num<10{
        return num
    }
    var tmp int
    for num>0{
        tmp=tmp+num%10
        num/=10
    }
    return addDigits(tmp)
}
