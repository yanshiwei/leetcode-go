func trailingZeroes(n int) int {
    //https://leetcode-cn.com/problems/factorial-trailing-zeroes/solution/c-shu-xue-xiang-xi-tui-dao-by-zeroac/
    var res int
    for n>0{
        res+=n/5
        n=n/5
    }
    return res
}
