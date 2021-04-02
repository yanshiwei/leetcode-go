func maxSideLength(mat [][]int, threshold int) int {
    /*
    参考：https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/solution/qing-xi-tu-jie-mo-neng-de-qian-zhui-he-by-hlxing/
    遍历所有可能的正方形区域，具体的算法流程是：1.考虑正方形的边长从 1 到 Min(M,N)Min(M,N)（M 为矩阵长度，N 为矩阵宽度）2.考虑正方形右下角的坐标从 (0, 0) 到 (M, N) 3.判断正方形是否存在（可能会超出边界，通过左上角坐标判断），如果存在则验证该正方形区域的元素总和。
    1.前缀和来计算每个正方形和，二维前缀和 preSum[i][j]：从 (第0, 第0) 到 (第i, 第j) 内元素的总和
    2.则可求出preSum[i][j] = mat[i-1][j-1] + preSum[i - 1][j] + preSum[i][j - 1] - preSum[i - 1][j - 1]
    3.从正方形右下角出发，已知某节点(i-1,j-1)假如作为正方形右下角坐标为 (i-1, j-1)，设此正方形边长 k，则正方形和：
    preSum[i][j]-preSum[i-k][j]-preSum[i][j-k]+preSum[i-k][j-k]
    */

    var m=len(mat)
    var n=len(mat[0])
    var preSum=make([][]int,m+1)
    for i:=range preSum{
        preSum[i]=make([]int,n+1)
    }
    for i:=1;i<=m;i++{
        for j:=1;j<=n;j++{
            preSum[i][j]=mat[i-1][j-1]+preSum[i-1][j]+preSum[i][j-1]-preSum[i-1][j-1]
        }
    }
    var res int
    for k:=1;k<=min(m,n);k++{
        //不断尝试正方形边长
        for i:=1;i<=m;i++{
            for j:=1;j<=n;j++{
                //边界判断
                if i-k<0||j-k<0{
                    continue
                }
                tmp:=preSum[i][j]-preSum[i-k][j]-preSum[i][j-k]+preSum[i-k][j-k]
                if tmp<=threshold{
                    res=max(res,k)
                }
            }
        }
    }
    return res
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
