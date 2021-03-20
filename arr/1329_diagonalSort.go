func diagonalSort(mat [][]int) [][]int {
    /*
    利用左对角线元素 坐标 i-j 相等的特性（右对角线元素 i+j 相等）
    把同一斜边的元素放到一个数组里排序
    再放回去
    */
    var m=len(mat)
    var n=len(mat[0])
    var idxM=make(map[int][]int)
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            idxM[i-j]=append(idxM[i-j],mat[i][j])
        }
    }
    for _,v:=range idxM{
        sort.Ints(v)
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            head:=idxM[i-j][0]
            mat[i][j]=head
            idxM[i-j]=idxM[i-j][1:]
        }
    }
    return mat
}
