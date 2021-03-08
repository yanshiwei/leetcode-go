func maxScoreSightseeingPair(values []int) int {
    /*
    我们考虑从前往后遍历 j来统计答案，对于每个观光景点 j而言，我们需要遍历 [0,j−1] 的观光景点 i 来计算组成观光组合 (i,j) 得分的最大值 cnt j ​ 来作为第 j个观光景点的值，那么最后的答案无疑就是所有观光景点值的最大值，即 max(cntj),其中j=0,1,...n-1。但是遍历 j需要O(n) 的时间复杂度，遍历 [0,j-1]的观光景点 ii也需要 O(n)的时间复杂度，因此该方法总复杂度为 O(n^2)，不能通过所有测试用例，我们需要进一步优化时间复杂度。
     我们回过头来看得分公式，我们可以将其拆分成values[i]+i 和 values[j]−j 两部分，这样对于统计景点 j 答案的时候，由于values[j]−j 是固定不变的，因此最大化values[i]+i+values[j]−j 的值其实就等价于求 [0,j−1] 中 values[i]+i 的最大值mx，景点 j 的答案即为 mx+values[j]−j 。而 mx 的值我们只要从前往后遍历 j 的时候同时维护即可，这样每次遍历到景点 j 的时候，寻找使得得分最大的 i 就能从 O(n)降至 O(1) 的时间复杂度，总时间复杂度就能从 O(n^2) 降至 O(n)
    */
    var res int
    var mx=values[0]+0//[0-j-1]中values[i]+i最大值
    for j:=1;j<len(values);j++{
        res=max(res,mx+values[j]-j)//此时mx是0-j-1]中values[i]+i最大值
        mx=max(mx,values[j]+j)//边遍历j时边维护
    }
    return res
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
