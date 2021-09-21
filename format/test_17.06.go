func numberOf2sInRange(n int) int {
    //seen https://leetcode-cn.com/problems/number-of-digit-one/
    if n<2{
        return 0
    }
    if n<12{
        return 1
    }
    if n<20{
        return 2
    }
    var idx int=1
    var sum int
    var befor int
    for n>0{
        tmp:=n%10
        n/=10
        if tmp>2{
            sum+=(n+1)*idx
        }else if tmp==2{
            sum+=(n*idx)+befor+1
        }else{
            sum+=n*idx
        }
        befor+=tmp*idx
        idx*=10
    }
    return sum
}
