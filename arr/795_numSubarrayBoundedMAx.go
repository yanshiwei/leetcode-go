func numSubarrayBoundedMax(A []int, L int, R int) int {
    /*
    区间大于等于L小于等于R的子数组个数=区间小于等于R的子数组个数-区间小于L的子数组个数
    也=区间小于等于R的子数组个数-区间小于等于L-1的子数组个数；
按照这个思想，遍历两边数组，算出来一减。
    如A=【2，1，4，3】，bound=4
    当i=0 A[0]<=bound 当前子数组为【1】数目1，累计数目1
    当i=1 A[1]<=bound 当前子数组为【1】，【2，1】数目2，累计数目1+2=3
    当i=2 A[2]<=bound 当前子数组为【4】，【1，4】，【2，1，4】数目3，累计数目3+3=6
    当i=3 A[0]<=bound 当前子数组为【3】，【4，3】，【1，4，3】，【2，1，4，3】数目4，累计数目6+4=10
    */
    return numberNotBiggerBound(A,R)-numberNotBiggerBound(A,L-1)
}
//计算所有元素都小于等于bound的子数组数量
func numberNotBiggerBound(A []int,bound int)int{
    var res int
    var cur int
    for i:=range A{
        if A[i]<=bound{
            cur++
        }else{
            //重新置0
            cur=0
        }
        res+=cur
    }
    return res
}
