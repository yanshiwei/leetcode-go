func rankTeams(votes []string) string {
    //var m=len(votes)
    var n=len(votes[0])
    var fre=make([][]int,26)
    for i:=range fre{
        fre[i]=make([]int,n)
    }
    var record=make(map[int]struct{})
    for i:=range votes[0]{
        record[int(votes[0][i]-'A')]=struct{}{}
    }
    //以示例1为例：记录A，B, C分别在每个排位出现了几次
    for i:=range votes{
        vote:=votes[i]
        for j:=0;j<n;j++{
            fre[vote[j]-'A'][j]++
        }
    }
    var pos=make([]Info,26)
    for i:=range fre{
        pos[i]=Info{
            Arr:fre[i],
            Idx:i,
        }
    }
    sort.Slice(pos,func(i,j int)bool{
        //挨个比较他们在排位中出现的次数
        p1:=pos[i]
        p2:=pos[j]
        for i:=0;i<n;i++{
            if p1.Arr[i]!=p2.Arr[i]{
                return p1.Arr[i]>p2.Arr[i]
            }
        }
        //第二种情况，比较他们的字母顺序
        return p1.Idx<p2.Idx
    })
    //fmt.Println(pos)
    var res string
    for i:=range pos{
        if _,ok:=record[pos[i].Idx];ok==false {
            continue
        }
        res+=string(pos[i].Idx+'A')
    }
    return res
}

type Info struct{
    Arr []int
    Idx int
    Exist bool
}
