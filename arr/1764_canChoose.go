func canChoose(groups [][]int, nums []int) bool {
    var n=len(groups)
    //nums 中选出 n 个 不相交 的子数组
    //判断数组group能否在nums中找到
    isFind:=func(group []int,innerNums[]int)bool{
        if len(group)>len(innerNums){
            return false
        }
        for i:=range group{
            if group[i]!=innerNums[i]{
                return false
            }
        }
        return true
    }
    //在nums中按顺序找groups[j], 找到后就在剩余数组中继续查找其他数组
    //将groups数组全部找到, 返回ture 
    var i,j int
    for i=0;i<len(nums)&&j<n;i++{
        if isFind(groups[j],nums[i:]){
            //找到nums就更新下标
            i+=len(groups[j])-1//-1是因为循环里还有i++
            j++
        }
    }
    if j==len(groups){
        return true
    }
    return false
}
