func isPossible(nums []int) bool {
    var FreMap=make(map[int]int)//建立数和其在nums出现次数的映射
    var CanPutTailMap=make(map[int]int)//数可以出现在某连续序列的后面，value是这个数可以出现的次数
    for i:=range nums {
        FreMap[nums[i]]++
    }
    for i:=range nums {
        if FreMap[nums[i]]==0 {
            continue//
        }else if CanPutTailMap[nums[i]]>0 {
            //说明当前这个数可以加到某个连续子序列的末尾，我们将其加上去
            CanPutTailMap[nums[i]]--
            FreMap[nums[i]]--
            CanPutTailMap[nums[i]+1]++
            //然后将下一个需要的连续数（比当前数大1）在CanPutTailMap中的映射值加1
        }else if FreMap[nums[i]+1]>0&&FreMap[nums[i]+2]>0{
            //如果当前数不能连接到其他子序列后面，那么我们看看这个数是否可以成为新的子序列的起点，如果这个数后面两个连续数在freq中的映射值都大于0的话，说明当前数可以作为新的子序列的起点
            FreMap[nums[i]]--
            FreMap[nums[i]+1]--
            FreMap[nums[i]+2]--
            CanPutTailMap[nums[i]+3]++
            //然后将这三个数下一个需要的连续数在need中的映射值加1
        }else {
            return false
        }
    }
    return true
}
