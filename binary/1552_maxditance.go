func maxDistance(position []int, m int) int {
    /*
    二分法
    0.最小磁力指的是这 m 个球中相邻两球距离的最小值的结论。对于i<j<k 三个位置的球，最小磁力一定是j−i 和k−j 的较小值，而不是跨越了位置 j的 i 和 k 的差值 k−i。
    1.比如说我们假设最小距离是3 那么两球间的距离至少是3不会低于3，否则3就不可能是最小距离。
    如果在最小距离是3的情况可以放置，那么我们就尝试最小距离为4的放置方案,
    然后尝试最小间隔为5
    在可放置完成与不可放置完成之间的距离就是答案。
    这是一个区间划分的问题 可以使用二分解决
    2.问题的关机是如何判断当前最小距离是否合法。即m个球最小距离是minDistance下 长度为n的position数组能否放下m个球。
    从贪心的角度考虑，第一个小球放置的篮子一定是position 最小的篮子，即排序后的第一个篮子。那么为了满足上述条件，第二个小球放置的位置一定要大于等于position[0]+minDistance，接下来同理。因此我们从前往后扫position 数组，看在当前答案 minDistance 下我们最多能在篮子里放多少个小球，我们记这个数量为cnt，如果cnt 大于等于 m，那么说明当前答案下我们的贪心策略能放下 m 个小球且它们间距均大于等于 minDistance ，为合法的答案，否则不合法
    */
    sort.Ints(position)
    var left=1
    var right=position[len(position)-1]-position[0]
    var ans =-1
    for left<=right{
        var mid=left+(right-left)/2
        if check(position,mid,m){
            //如果该距离合理，则尝试更大的距离
            ans=mid
            left=mid+1
        }else{
            //如果该距离不合理，则尝试更小的距离
            right=mid-1
        }
    }
    return ans
}
//m个球最小距离是minDistance下 长度为n的position数组能否放下m个球
func check(position[]int, minDistance,m int)bool{
    //第一个球肯定在起点
    var pos=position[0]
    var cnt=1
    for i:=1;i<len(position);i++{
        if position[i]-pos>=minDistance{
            //可以放下一个球
            pos=position[i]
            cnt++
        }
    }
    if m<=cnt{
        return true
    }
    return false
}
