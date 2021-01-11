func nextPermutation(nums []int)  {
    /*
    如果存在更大一个排列则给出
    如果不存在，则是从开始到结束是降序如321，直接反序即可
    故从后到前，如果存在一个元素a[i]<a[i+1]则存在下一个更大排列
    然后查找nums[i]右边降序结构中，比nums[i]大的最小值nums[k]，交换
    后将nums[i]后的数字升序排列即可
    证明过程：
    字典序下一个排列其实就是让当前的排序变大一点点，幅度为最小的情况。想要变大，就一定是找到一个顺序排列的子排列，将大的移动到前面，这样就会变大。
    那么为了变大一丢丢，就应该找到当前排列最后面的一个小顺序，然后将后面大一点点的数移到前面。
具体算法：
    1/ 从后向前遍历数组，找到第一个nums[i] > nums[i - 1]。这个操作保证了区间[i,end)上的元素都是递减排列的。
    2/ 再从后遍历区间[i, end)，找到第一个比nums[i - 1]大的元素，交换这两个元素。
    3/ 这时[i,end)区间上的元素依然是递减排列的，因此可以直接翻转这个区间上的元素，这样就达到了只增大一点点的目的
    总结：
    最简单的思路，就是将其当做一个十进制的数去理解，找出比当前十进制数大的下一位，比如1234，下一个就是1243。

第一步：从最后一位元素找到要替换数组元素的下标swapIndex。替换满足nums[i]>nums[i-1]，此时swapIndex=i-1。比如：[1,2,3,4]， 应该返回的就是3这个下标。如果是4321，则应该没有满足此项要求的下标。则认为此数组降序排列，可以直接对该数组依次首尾交换元素即可。程序结束。

第二步： 从替换数组元素的下一位（swapIndex+1）开始寻找，比替换元素大的最小的元素，并将两者进行替换。比如[1,2,3,4],返回的是3的下标，那么从4这个元素开始往下找剩下的元素，从剩下的元素中找到一个最小元素且该元素比3大。然后在交换这两个元素。

第三步：从swapIndex+1开始，对数组剩余的元素进行翻转，（此时就是从swapIndex+1后的元素就是升序了）。
注意：第二步找到替换的元素时，会有重复的数字，因此要注意等号的判断。
    */
    if len(nums)<2{
        return
    }
    var isReverse bool
    for i:=len(nums)-2;i>=0;i--{
        if nums[i]>=nums[i+1]{continue}
        isReverse=true
        j:=len(nums)-1
        for nums[j]<=nums[i]{j--}
        nums[i],nums[j]=nums[j],nums[i]
        sort.Ints(nums[i+1:])
        break
    }
    if isReverse ==false {
        //已经是全部降序
        sort.Sort(sort.Reverse(IntSlice(nums)))
    }
    return
}

type IntSlice []int
 
 
 
func (s IntSlice) Len() int { return len(s) }
 
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
 
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
