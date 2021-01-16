func jump(nums []int) int {
    /*
    贪心算法
    参考：
    https://leetcode-cn.com/problems/jump-game-ii/solution/tiao-yue-you-xi-ii-by-leetcode-solution/
    https://leetcode-cn.com/problems/jump-game-ii/solution/tan-xin-shi-jian-fu-za-du-onkong-jian-fu-55fq/
例如，对于数组 [2,3,1,2,4,2,3]，初始位置是下标 0，从下标 0 出发，最远可到达下标 2。下标 0 可到达的位置中，可以达到1也可以达到2，下标 1 的值是 3（最终到位置4）而下标2的值是1（最终到位置3），从下标 1 出发可以达到更远的位置，因此第一步到达下标 1。
    从下标1出发，最远距离时3，故可以到下标2 3 4
    下标2的值是1（最终到位置3）；下标3的值是2（最终到大位置5）；下标4的值时4（最终到达位置8），因此选择下标4
    以此类推
    总结：
    0/ maxPos代表某次跳跃的跳跃终点（下次跳跃长度最远的点作为该次跳跃的终点）,lastMaxPos代表上一个跳跃的终点位置。由于数组都是正数，故在遍历最后一个元素前，一定能到达最后一个为主，故不必访问最后一个元素。
    1每次跳跃取跳跃区间内的下次跳跃长度最远的点作为该次跳跃的终点maxPos=max(maxPos,i+nums[i]) ；
    2 遍历i-len(nums-1),不断更新某次跳跃的跳跃终点走到上次跳跃的区间终点时,更新下次跳跃的区间范围,跳跃次数+1
    即
    定义步数step=0，能达到的最远位置max_bound=0，和上一步到达的边界end=0。
    遍历数组，遍历范围[0,n−1)：
        所能达到的最远位置max_bound=max(max_bound,nums[i]+i)，表示上一最远位置和当前索引i和索引i上的步数之和中的较大者。
        如果索引i到达了上一步的边界end，即i==end，则：
        更新边界end，令end等于新的最远边界dmax_bound
b
​	
 ound
令步数stepstep加一
返回step
    */
    if len(nums)<=1{
        return 0
    }
    var lastMaxPos=nums[0] 
    var maxPos =lastMaxPos
    var steps =1
    for i:=0;i<len(nums)-1;i++{
        maxPos=max(maxPos,i+nums[i])
        if i==lastMaxPos{
            //走到了上次跳跃的区间终点
            lastMaxPos=maxPos
            steps++
        }
    }
    return steps
}
func max(x,y int)int{
    if x<y{
        return y
    }
    return x
}
