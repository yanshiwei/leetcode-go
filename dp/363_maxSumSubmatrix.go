func maxSumSubmatrix(matrix [][]int, k int) int {
    /*
    参考：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/solution/czhuang-tai-ya-suo-dong-tai-gui-hua-by-cainhsu/
    https://leetcode-cn.com/problems/max-submatrix-lcci/solution/zhe-yao-cong-zui-da-zi-xu-he-shuo-qi-you-jian-dao-/
    https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/solution/javascriptjie-fa-can-kao-zui-da-zi-ju-zhen-de-si-l/
    53.最大子序列和：https://leetcode-cn.com/problems/maximum-subarray/submissions/
    */

    var row=len(matrix)
    var col=len(matrix[0])
    var storage=make([]int,col)//将matrix[i...j][k]这一列数据放入storage[k]中
    var res=INTMIN
    var lastDp int//相当于dp[i-1]
    //记录当前i~j行组成大矩阵的每一列的和，将二维转化为一维
    for i:=0;i<row;i++ {
        //以i为上边，从上而下扫描
        for t:=0;t<col;t++ {
            storage[t]=0
        } //每次更换子矩形上边，就要清空b，重新计算每列的和
        for j:=i;j<row;j++{
            //子矩阵的下边，从i到col-1，不断增加子矩阵的高
            //执行一次就相当于求一次最大子序列和
            lastDp=0;//从头开始求dp
            for kk:=0;kk<col;kk++{
                storage[kk]+=matrix[j][kk]
            }
            //遍历所有，因为包含K的限制 所以不直接DP
            for m:=0;m<col;m++ {
                lastDp=0
                for n:=m;n<col;n++ {
                    lastDp+=storage[n]
                    if lastDp<=k&& lastDp>res{
                        res=lastDp
                    }
                }
            }
        }
    }
    return res
}
const INTMAX=int(^uint(0) >> 1)
const INTMIN=^INTMAX
