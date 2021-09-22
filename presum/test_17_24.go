func getMaxMatrix(matrix [][]int) []int {
    //prefix[x][y] 表示从 [0, 0] 到[x, y]的子矩阵和。
    //则prefix[x][y]=m[x][y]+prefix[x-1][y]+prefix[x][y-1]-prefix[x-1][y-1]
    var res=make([]int,4)
    var n=len(matrix)
    if n<1{
        return res
    }
    var m=len(matrix[0])
    if m<1{
        return res
    }
    var preSum=make([][]int,n+1)
    for i:=range preSum{
        preSum[i]=make([]int,m+1)
    }
    for i:=1;i<=n;i++{
        for j:=1;j<=m;j++{
            preSum[i][j]=preSum[i][j-1]+preSum[i-1][j]-preSum[i-1][j-1]+matrix[i-1][j-1]
        }
    }
    var globalMAx=IntMin
    //先固定上下两条边
    for top:=0;top<n;top++{
        for bottom:=top;bottom<n;bottom++{
            localMax:=0
            left:=0
            //然后从左往右一遍扫描找最大子序和
            for right:=0;right<m;right++{
                //利用presum快速求出localMax
                localMax=preSum[bottom+1][right+1]+preSum[top][left]-preSum[top][right+1]-preSum[bottom+1][left]
                if localMax>globalMAx{
                    globalMAx=localMax
                    res[0]=top
                    res[1]=left
                    res[2]=bottom
                    res[3]=right
                }
                //如果不满0，前面都舍弃，从新开始计算，left更新到right+1，right下一轮递增之后left==right
                if localMax<0{
                    localMax=0
                    left=right+1
                }
            }
        }
    }
    return res
}
const IntMax=int(^uint(0)>>1)
const IntMin=^IntMax
