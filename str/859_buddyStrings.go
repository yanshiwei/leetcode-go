func buddyStrings(s string, goal string) bool {
    var n=len(s)
    var m=len(goal)
    if n!=m{
        return false
    }
    var cnt1=make(map[byte]int)
    var cnt2=make(map[byte]int)
    var diffSum int
    for i:=0;i<n;i++{
        cnt1[s[i]]++
        cnt2[goal[i]]++
        if s[i]!=goal[i]{
            diffSum++
        }
    }
    var twoFlag=false
    for k,v:=range cnt1{
        if v!=cnt2[k]{
            // 频率不同，肯定不亲密
            return false
        }
        if v>1{
            //s 中有出现数量超过 2 的字符（
            twoFlag=true
        }
    }
    // s gola字符不同的数量为2，直接交换，肯定亲密
    if diffSum==2{
        return true
    }else{
    // s gola字符不同的数量为0时，还要求同时 s 中有出现数量超过 2 的字符（能够相互交换）
        for k,v:=range cnt1{
            if v!=cnt2[k]{
                // 频率不同，肯定不亲密
                return false
            }
            if v>1{
                //s 中有出现数量超过 2 的字符（
                twoFlag=true
            }
        }
        if diffSum==0&&twoFlag{
            return true
        }
        return false
    }
}
