func countSquares(matrix [][]int) int {
    /*
    动态规划：
    参考：https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/solution/tong-ji-quan-wei-1-de-zheng-fang-xing-zi-ju-zhen-2/
    证明：
    1.固定 f[i][j] 的值。若对于位置 (i, j) 有 f[i][j] = 4，我们将以 (i, j) 为右下角、边长为 4 的正方形涂上色，可以发现其左侧位置 (i, j - 1)，上方位置 (i - 1, j) 和左上位置 (i - 1, j - 1) 均可以作为一个边长为 4 - 1 = 3 的正方形的右下角。也就是说，这些位置的的 f 值至少为 3，即：
    min(f[i][j−1],f[i−1][j],f[i−1][j−1])≥f[i][j]−1
    2.也可以固定 f[i][j] 相邻位置的值，得到另外的限制条件。假设 f[i][j - 1]，f[i - 1][j] 和 f[i - 1][j - 1] 中的最小值为 3，也就是说，(i, j - 1)，(i - 1, j) 和 (i - 1, j - 1) 均可以作为一个边长为 3 的正方形的右下角。我们将这些边长为 3 的正方形依次涂上色，可以发现，如果位置 (i, j) 的元素为 1，那么它可以作为一个边长为 4 的正方形的右下角，f 值至少为 4，即：
    f[i][j]≥min(f[i][j−1],f[i−1][j],f[i−1][j−1])+1
    将其与上一个不等式联立，可以得到：
    f[i][j]=min(f[i][j−1],f[i−1][j],f[i−1][j−1])+1
    此处f[i][j]记录的是：
    以此点（i，j）为右下角的正方形的个数
    转移方程：
    f[i][j] = min(f[i - 1][j - 1],min(f[i - 1][j],f[i][j - 1])) + 1;
    边界条件：
    if matrix[i][j]=0 f=0
    if i==0||j==0 f[i][j]=matrix[i][j]
    */
    var m=len(matrix)
    var n=len(matrix[0])
    var f=make([][]int,m)
    for i:=range f{
        f[i]=make([]int,n)
    }
    var res int
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if matrix[i][j]==0{
                f[i][j]=0
            }else if i==0||j==0{
                f[i][j]=matrix[i][j]
            }else{
                //左上角 左侧 上侧
                f[i][j]=min(f[i-1][j-1],min(f[i-1][j],f[i][j-1]))+1
            }
            res+=f[i][j]
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
