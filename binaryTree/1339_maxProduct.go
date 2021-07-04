/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var sum int
 var best int
func maxProduct(root *TreeNode) int {
    if root==nil{
        return 0
    }
    sum=0
    best=0
    dfsForAllSum(root)
    dfsForEachMax(root)
    //fmt.Println(sum,best)
    res:=int64(best)*int64(sum-best)%1000000007
    return int(res)
}

func dfsForAllSum(root *TreeNode){
    if root==nil {
        return
    }
    sum+=root.Val
    dfsForAllSum(root.Left)
    dfsForAllSum(root.Right)
}
//某节点作为根节点的子树的和
func dfsForEachMax(root *TreeNode)int{
    if root==nil{
        return 0
    }
    curTreeSum:=dfsForEachMax(root.Left)+dfsForEachMax(root.Right)+root.Val
    if abs(curTreeSum*2-sum)<abs(best*2-sum){
        best=curTreeSum
    }
    return curTreeSum
}

func abs(x int)int {
    if x <0{
        return -1*x
    }
    return x
}

