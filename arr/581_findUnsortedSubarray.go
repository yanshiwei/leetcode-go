func findUnsortedSubarray(nums []int) int {
    var tmp []int
    for i:=range nums {
        tmp=append(tmp,nums[i])
    }
    sort.Ints(nums)
    var left=len(nums)
    var right=0
    for i:=range nums {
        if nums[i]!=tmp[i] {
            //最左边和最右边不匹配的元素
            left=min(left,i)
            right=max(right,i)
        }
    }
    if right-left<0 {
        return 0
    }
    return right-left+1
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
