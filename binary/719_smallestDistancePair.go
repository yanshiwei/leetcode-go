func smallestDistancePair(nums []int, k int) int {
    /*
    if len(nums)<2{
        return 0
    }
    sort.Ints(nums)
    var minHeap=&IntHeap{}
    heap.Init(minHeap)
    //排序后，(i, j + 1) 的距离不会小于 (i, j)，因此如果 (i, j) 还在堆中，我们没有必要把 (i, j + 1) 放入堆中故可以先只把
    //0-1，1-2... i-i+1... len(nums)-2-len(nums)-1放入堆
    //而不用0-1 0-2 0-3...
    //build pair
    for i:=0;i<len(nums)-1;i++ {
        info:=Info{X:i,Y:i+1,Dis:nums[i+1]-nums[i]}
        heap.Push(minHeap,info)
    }
    fmt.Printf("before%+v",minHeap)
    var kthDis int
    for k>0 {
        k--
        info:=heap.Pop(minHeap).(Info)
        kthDis=info.Dis
        if info.Y+1<len(nums){
            infoNext:=Info{X:info.X,Y:info.Y+1,Dis:nums[info.Y+1]-nums[info.X]}
            heap.Push(minHeap,infoNext)
        }
    }
    return kthDis
    */
    //time out above
    /*
   1. 对(A,B)的距离的值进行二分搜索，对每个距离判断小于当前距离的个数n，如果n<k 舍弃左半集合，否则，舍弃右半集合
   2. 在计算小于当前距离的个数n时，使用双指针的方法，见caculationCnt函数
    */
    //二分搜索+双指针
    sort.Ints(nums)
    var leftDis =0//min dis
    var rightDis=nums[len(nums)-1]-nums[0]//max dis
    //第k小的距离差必然在 [min_distance, max_distance] 之间
    //通过二分搜索确定距离差
    for leftDis<rightDis {
        var midDis=(rightDis+leftDis)/2
        // 淘汰策略
        // 对于mid而言
        //若小于mid的距离差总数 >= k，则距离差应落在 [low, mid] 之间
        //若小于mid的距离差总数 < k，则距离差应落在 [mid+1, high] 之间
        if caculationCnt(nums,midDis)>=k{
            rightDis=midDis
        }else {
            leftDis=midDis+1
        }
    }
    return leftDis
}

func caculationCnt(nums[]int,targetDis int)int {
    //由于数组已有序，所以我们只需要统计差值在target内的数量即可
    //大于target的我们可以直接跳过，以此来减少计算次数
    //依然使用动态窗口机制，我们每次计算至差值 <= target
    //则窗口向右滑动时，两侧元素差值 > target，我们可以直接将左侧元素剔除
    var leftIdx,cnt int
    for rightIdx:=1;rightIdx<len(nums);rightIdx++ {
        for nums[rightIdx]-nums[leftIdx]>targetDis {
            leftIdx+=1
        }
        cnt+=rightIdx-leftIdx
    }
    return cnt
}
