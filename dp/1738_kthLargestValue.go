func kthLargestValue(matrix [][]int, k int) int {
    /*
    动态规划
    堆排序
    参考https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value/solution/tu-jie-dong-tai-gui-hua-by-yexiso-ln70/
    坐标[i,j]左上方的所有元素的异或值，这里用f[i][j]来表示
    f[i][j] = f[i - 1][j - 1] ^ f[i][j - 1] ^ f[i - 1][j] ^ matrix[i][j]
    代码中的数组f下标从1开始
    */
    var m=len(matrix)
    var n=len(matrix[0])
    var f=make([][]int,m+1)
    for i:=range f{
        f[i]=make([]int,n+1)
    }
    var minHeap=&IntHeap{}
    heap.Init(minHeap)
    for i:=1;i<=m;i++{
        for j:=1;j<=n;j++{
            f[i][j]=f[i-1][j-1]^f[i-1][j]^f[i][j-1]^matrix[i-1][j-1]
            heap.Push(minHeap,f[i][j])
        }
    }
    //fmt.Println(minHeap)
    for {
        k--
        if k<1{
            break
        }
        heap.Pop(minHeap)
    }
    return minHeap.Top()
}

type IntHeap []int
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int)bool{return h[i]>h[j]}//min
func (h *IntHeap)Push(x interface{}){
    *h=append(*h,x.(int))
}
func (h IntHeap)Top()int{
    return h[0]
}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
