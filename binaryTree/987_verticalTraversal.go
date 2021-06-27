/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var minHeap *IntHeap
func verticalTraversal(root *TreeNode) [][]int {
    var res [][]int
    if root==nil{
        return res
    }
    minHeap=&IntHeap{}
    heap.Init(minHeap)
    dfs(root,0,0)
    // 遍历优先级队列然后构建res
    //fmt.Println(minHeap)
    var preY=INTMIN
    for minHeap.Len()>0{
        cur:=minHeap.Top()
        v:=cur.Val
        y:=cur.Col
        if preY==INTMIN||preY!=y{
            //first or diff
            tmp:=make([]int,0)
            res=append(res,tmp)
            preY=y
        }
        res[len(res)-1]=append(res[len(res)-1],v)
        heap.Pop(minHeap)
    }
    return res
}

func dfs(cur *TreeNode,x,y int){
    if cur!=nil{
        //插入堆
        heap.Push(minHeap,Info{Val:cur.Val,Row:x,Col:y})
        dfs(cur.Left,x+1,y-1)
        dfs(cur.Right,x+1,y+1)
    }
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
type IntHeap []Info
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){
    h[i],h[j]=h[j],h[i]
}
func (h IntHeap)Less(i,j int)bool{
    if h[i].Col!=h[j].Col{
        //按照col从小到大
        return h[i].Col<h[j].Col
    }else if h[i].Row!=h[j].Row{
        //按照row从小到大
        return h[i].Row<h[j].Row
    }else{
        //按照val从小到大
        return h[i].Val<h[j].Val
    }
}
func (h IntHeap)Top()Info{
    return h[0]
}
func (h *IntHeap)Push(x interface{}){
    *h=append(*h,x.(Info))
}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[:n-1]
    return x
}
type Info struct{
    Val int
    Row int
    Col int
}
