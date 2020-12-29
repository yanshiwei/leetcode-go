func kthSmallestPrimeFraction(A []int, K int) []int {
    /*
    思路:
    under(x) 用于求解小于 x 的分数数量，这是一个关于 x 的单调增函数，因此可以使用二分查找求解。
    算法:
    使用二分查找找出一个 x，使得小于 x 的分数恰好有 K 个，并且记录其中最大的一个分数。
    */
    var res []int
    if len(A)<2 {
        return res
    }
    //A already sort and A[0] must be 1
    var leftDis=0.0
    var rightDis=1.0//A[0]
    res=append(res,0)
    res=append(res,1)
    for rightDis-leftDis>1e-9 {
        //注意浮点型比较法
        midDis:=(leftDis+rightDis)/2
        cacRes:=caculateCnt(A,midDis)//小于midDis的个数 及分数值最大的分子和分母
        if cacRes[0]<K{
            //小于 mi 的分数数量小于 K，更新区间为 [mi+1, right]
            leftDis=midDis
        }else {
            //小于 mi 的分数数量不小于 K，更新区间为 [left, mi]
            rightDis=midDis
            res[0]=cacRes[1]
            res[1]=cacRes[2]
        }
    }
    return res
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
//二分实现找lowerBound
//使用滑动窗口的方法：对于每个 primes[j]，找出最大的 i 使得
// primes[i] / primes[j] < x。
//随着 j （和 primes[j]）的增加， i 也会随之增加。
func caculateCnt(primes []int, target float64)[]int {
    // Returns {count, numerator, denominator}
    var numerator=0
    var denominator=1
    var count=0
    var i =-1
    for j:=1;j<len(primes);j++ {
        for float64(primes[i+1])<float64(primes[j])*target {
            //p[i]/p[j]<x
            i++
        }
        count+=i+1//before i primes[i-] is smaller and primes[j] bigger so prime[i]/prime[j] must less than x count is current
        fmt.Println(i,j,count,target)
        if i>=0&&numerator*primes[j]<denominator*primes[i]{
            //p[i]/p[j]>numerator/denominator
            //find bigger numerator/denominator
            numerator=primes[i]
            denominator=primes[j]
        }
    }
    return []int{count,numerator,denominator}
}
