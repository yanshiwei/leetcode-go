func rangeBitwiseAnd(left int, right int) int {
    //https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--41/
    var zeros int// 记录我们右移的次数，也就是最低位 0 的个数
    for left<right{
        zeros++
        left>>=1
        right>>=1
    }
    return left<<zeros//将 0 的个数空出来
}
