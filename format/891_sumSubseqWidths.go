func sumSubseqWidths(A []int) int {
    /*
    排完序之后，我们可以观察到：一个元素是最大元素：当且仅当它是被选取元素中最右边的一个；一个元素是最小元素，当且仅当它是被选取元素中最左边的一个。所以说，假设排序后一个元素左边有left个元素，右边有right个元素，那么这个元素作为最小值出现的子序列一共有2^{right}
个（右边的每个元素可以选取或不选取）；而作为最大值出现的子序列一共有2^{left}个。因此，元素A[i]对最后的总和的贡献就等于
最小值：-A[i]*2^right
最大值：A[i]*2^left
A[i]*2^left+（-A[i]*2^right）=(2^i-2^n-i-1)*A[i]
所有元素对结果影响：
sum((2^i-2^n-i-1)*A[i]),i from 0 to n-1
    */
    var mod int64=1000000007
    var n=len(A)
    sort.Ints(A)
    var pow2=make([]int64,n)
    pow2[0]=1
    for i:=1;i<n;i++{
        pow2[i]=pow2[i-1]*2%mod//防止过大
    }
    var res int64
    for i:=0;i<n;i++{
        res=(res+(pow2[i]-pow2[n-i-1])*int64(A[i]))%mod
    }
    return int(res)
}
