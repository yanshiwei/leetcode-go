/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var res[]int
func preorder(root *Node) []int {
    res=nil
    preorderInner(root)
    return res
}

func preorderInner(root *Node){
    if root==nil{
        return
    }else{
        res=append(res,root.Val)
    }
    for i:=range root.Children{
        preorderInner(root.Children[i])
    }
}
