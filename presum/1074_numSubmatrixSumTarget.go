func numSubmatrixSumTarget(matrix [][]int, target int) int {
    /*
    前缀和
    prefix[x][y] 表示从 [0, 0] 到[x, y]的子矩阵和。
    则prefix[x][y]=m[x][y]+prefix[x-1][y]+prefix[x][y-1]-prefix[x-1][y-1]
    注意边界处理
    从[i,j]到[x,y]的任意子矩阵之和
    distance([i,j],[x,y])=prefix[x][y]-prefix[i-1][y]-prefix[x][j-1]+prefix[i-1][j-1]
    */
    var m=len(matrix)
    var n=len(matrix[0])
    var prefix=make([][]int,m)
    for i:=range prefix{
        prefix[i]=make([]int,n)
    }
    var res int
    for x:=0;x<m;x++{
        for y:=0;y<n;y++{
            if x==0&&y==0{
                //起始点
                prefix[x][y]=matrix[x][y]
            }else if x==0{
                //上侧边
                prefix[x][y]=matrix[x][y]+prefix[x][y-1]
            }else if y==0{
                //左侧边
                prefix[x][y]=matrix[x][y]+prefix[x-1][y]
            }else {
                prefix[x][y]=matrix[x][y]+prefix[x-1][y]+prefix[x][y-1]-prefix[x-1][y-1]
            }
            //distance([i,j],[x,y])=prefix[x][y]-prefix[i-1][y]-prefix[x][j-1]+prefix[i-1][j-1]=prefix[x][y]-a-b-c
            for i:=0;i<=x;i++{
                for j:=0;j<=y;j++{
                    var a,b,c int
                    if i==0||j==0{
                        c=0
                    }else{
                        c=prefix[i-1][j-1]
                    }
                    if i==0{
                        a=0
                    }else{
                        a=prefix[i-1][y]
                    }
                    if j==0{
                        b=0
                    }else{
                        b=prefix[x][j-1]
                    }
                    var tmpSum=prefix[x][y]-a-b+c
                    if tmpSum==target {
                        res++
                    }
                }
            }
        }
    }
    return res
}
