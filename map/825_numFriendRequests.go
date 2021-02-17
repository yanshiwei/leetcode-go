func numFriendRequests(ages []int) int {
    var count=make(map[int]int)
    for i:=range ages{
        count[ages[i]]++
    }
    var res int
    //对于每个元组 (ageA, countA)，(ageB, countB)，如果条件满足对应的年纪，那么久将 countA * countB 加入发好友请求的人数
    for ka,va:=range count{
        for kb,vb:=range count{
            if kb<=(ka/2+7){
                continue
            }
            if kb>ka{
                continue
            }
            if kb>100&&ka<100{
                continue
            }
            res+=va*vb
            if ka==kb{
                //当 ageA == ageB 的时候我们就数多了：我们只有 countA * (countA - 1) 对好友请求，因为你不能和自己发送请求
                res-=va
            }
        }
    }
    return res
}
