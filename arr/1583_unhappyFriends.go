func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
    var order =make([][]int,n)
    //order[i][j] 表示朋友 j在 i的朋友列表中的亲近程度下标。越亲近下标越小
    for i:=range order{
        order[i]=make([]int,n)
    }
    for i:=0;i<n;i++{
        for j:=0;j<n-1;j++{
            //注意j<n-1
            order[i][preferences[i][j]]=j
        }
    }
    var match=make([]int,n)
    //如果 x 和 y 配对，则有match[x]=y 以及 match[y]=x
    for i:=range pairs{
        pair:=pairs[i]
        p0:=pair[0]
        p1:=pair[1]
        match[p0]=p1
        match[p1]=p0
    }
    var res int
    for x:=0;x<n;x++{
        y:=match[x]
        idx:=order[x][y]
        //判断xy配对是不是合理
        for i:=0;i<idx;i++{
            //如果i<idx 说明存在xu亲近度超过xy
            u:=preferences[x][i]
            v:=match[u]
            //接下来只需要判断ux亲近度大于uv
            if order[u][x]<order[u][v]{
                res++
                //已经不合理，退出
                break
            }
        }
    }
    return res
}
