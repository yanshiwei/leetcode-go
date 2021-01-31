func minSubArrayLen(s int, nums []int) int {
    /*
    滑动窗口法
    1 变量start表示窗口的起始位置，变量end表示窗口的结束位置，区间[start,end]用于记录当前窗口中的元素，变量sum表示窗口中所有元素的总和，变量result表示符合题意的子数组长度
    2 滑动窗口的起始位置start=0，结束位置end=-1，表示，初始状态下窗口中没有元素，因为区间[0,-1]并不存在。result的初始值给定数组长度加1
    3 当窗口内序列和小于目标值时，扩大窗口即end++使得序列和变大
    4 当窗口内序列和不小于目标值时，按照题目要求的时最短子序列和，故没必要再扩大窗口，为了求最短长，反而可以尝试缩小窗口，即start++，窗口内的元素总和是否依旧大于等于目标值s=7，如果是，则继续缩小窗口左侧边界，如果不是则扩大窗口右侧边界，直到数组末尾。
    */
    if len(nums)<1 {
        return 0
    }
    var start=0
    var end =-1//表示窗口无数据
    var sum=0//窗口内所有数据和
    var minLen=len(nums)+1
    for start<len(nums){
        if end+1<len(nums)&&sum<s {
            //小于目标值，扩大窗口
            end++
            sum+=nums[end]
        }else {
            //大于等于目标值，满足要求，题目要求最短子序列，故没必要再扩大窗口，反而要尝试缩小窗口
            sum-=nums[start]
            start++
        }
        //窗口内值大于等于目标值的更新结果
        if sum>=s {
            minLen=min(minLen,end-start+1)
        }
    }
    if minLen==len(nums)+1 {
        //没找到
        minLen=0
    }
    return minLen
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
