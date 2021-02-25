func hasGroupsSizeX(deck []int) bool {
    /*
    统计每个数出现的次数，求所有次数的最大公约数即可
    注意两个数的最大公约数求法
    */
    var m=len(deck)
    if m<2{
        return false
    }
    var fre=make(map[int]int)
    for i:=range deck {
        if _,ok:=fre[deck[i]];ok==false {
            fre[deck[i]]=1
        }else{
            fre[deck[i]]++
        }
    }
    //出现次数的最大公约数
    var x int
    for _,cnt:=range fre {
        if cnt>0 {
            x=gcd(x,cnt)
            if x==1 {
                return false
            }
        }
    }
    return x>=2
}

func gcd(a,b int)int {
    if a<b {
        return gcd(b,a)
    }
    //a must >= b
    for b>0 {
        c:=a%b
        a=b
        b=c
    }
    return a
}
