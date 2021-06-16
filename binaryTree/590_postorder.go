/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var res []int
func postorder(root *Node) []int {
   res=nil
   postorderInner(root)
   return res 
}
func postorderInner(root *Node){
    if root==nil{
        return
    }else{
        for i:=range root.Children{
            postorderInner(root.Children[i])
        }
        res=append(res,root.Val)
    }
}
