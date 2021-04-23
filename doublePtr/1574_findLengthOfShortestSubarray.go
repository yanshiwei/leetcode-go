func findLengthOfShortestSubarray(arr []int) int {
    /*
    参考：
    https://leetcode-cn.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/solution/hua-hua-by-solituderain/
    把 数组 分为 A, B, C 三部分.
    A, C 为 非递减数组, B 为 中间部分.
    然后分情况讨论.
    图 画的是线性递增, 实际是其他的, 简化一下, 不影响结果.
    取其中最小的,
    case1.保留A, 删除 B + C
    case2.保留C, 删除 A + B
    case3.删掉 A一部分 + B + C一部分, 即 删掉 A[i:] + B + C[:j]

    其中 B, C可能为空,
    若B, C 都为空, 即本身就是非递减, 不用删
    若B为空, 可分为两段递增数组, 在case3里
    C为空, 删掉B.
    求 最小距离 d, 即 索引差.
    */
    var res int
    var n=len(arr)
    if n<2{
        return res
    }
    var left int
    for left+1<n&&arr[left]<=arr[left+1]{
        //非递减
        left++
    }
    if left==n-1{
        //B和C都为空，不删，都是非递减
        return res
    }
    var right=n-1
    for right-1>=0&&arr[right-1]<=arr[right]{
        //非递减
        right--
    }
    //左侧非递减left+1个，右侧非递减n-(right)个
    //若右侧为空，则最多移走n-left-1个；若左侧为空，则最多移走right个
    //初始化最大值
    res=min(n-left-1,right)
    var i int//i取值0-left代表A中一个点
    var j int=right//取值right-n-1代表C中一个点
    //left-right之间代表B
    for i<=left&&j<=n-1{
        //也是单调递减的，故谁小移谁就可以
        if arr[i]<=arr[j]{
            res=min(res,j-i-1)
            i++
        }else{
            j++
        }
    }
    return res
}

func min(x,y int)int {
    if x<=y {
        return x
    }
    return y
}
