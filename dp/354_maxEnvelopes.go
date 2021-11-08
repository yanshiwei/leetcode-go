func maxEnvelopes(envelopes [][]int) int {
    var n=len(envelopes)
    // 先按照宽做排序，越大的放在越前面，可以只看他的高度或者宽度，就简化成了一维数组
    sort.Slice(envelopes,func(i,j int)bool{
        return envelopes[i][0]>envelopes[j][0]
    })
    var dp=make([]int,n+1)//仅使用信封[0,i] (这里是区间的意思，表示前i+1 个信封)，且以第 i 个信封为顶端信封时的最大高度
    var res int
    // 依次尝试放置每个信封
    for i:=0;i<n;i++{
        // 遍历他之前的每个信封，看能放下他且最多层的个数
        var maxh int
        for j:=0;j<i;j++{
            // 判断是否可以放下当前的信封
            if envelopes[j][0]>envelopes[i][0]&&envelopes[j][1]>envelopes[i][1]{
                // 如果可以放下当前信封，看看是不是最大高度
                if maxh<dp[j]{
                    maxh=dp[j]
                }
            }
        }
        // 遍历一圈，找到最高，且能放下当前信封的maxh
        dp[i]=maxh+1
        // 判断当前信封高度是不是最高高度
        if res<dp[i]{
            res=dp[i]
        }
    }
    return res
}
