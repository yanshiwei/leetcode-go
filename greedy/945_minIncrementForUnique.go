func minIncrementForUnique(A []int) int {
    /*
    贪心算法在于每个子问题的局部最优解会指向全局最优解。
显然在对数组排序之后，可以通过保证数组的后一个元素，经过+1操作后比前面所有元素大即可,此时子问题的最优解会收敛于全局最优解
    */
    var m=len(A)
    sort.Ints(A)
    var res int
    for i:=1;i<m;i++{
        if A[i]<=A[i-1]{
            //A[i-1]在经过后面更新后 可能大于A[i]
            res+=A[i-1]-A[i]+1
            A[i]=A[i-1]+1
        }
    }
    return res
}
