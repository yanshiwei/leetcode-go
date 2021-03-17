func numEquivDominoPairs(dominoes [][]int) int {
    var res int
    var fre=make(map[int]int)
    for i:=range dominoes{
        one:=dominoes[i]
        var v int
        //1 <= dominoes[i][j] <= 9
        if one[0]>one[1]{
            v=one[1]*10+one[0]
        }else{
            v=one[0]*10+one[1]
        }
        //如果只有一个，则是0；如果有2个则是1；如果有3个则是2
        res+=fre[v]
        fre[v]++
    }
    return res
}
