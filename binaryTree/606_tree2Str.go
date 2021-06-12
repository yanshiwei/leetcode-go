/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res string
func tree2str(root *TreeNode) string {
    res=""
    preorder(root)
    if len(res)>0{
        //处理头尾两个多余括
        res=res[1:]
        res=res[0:len(res)-1]
    }
    return res
}

func preorder(root*TreeNode){
    if root==nil{
        return
    }
    res+="("
    res+=strconv.Itoa(root.Val)
    //当前访问的节点的左子树为空且右子树不为空，就在字符串中添加一个"()"
    if root.Left==nil&&root.Right!=nil{
        res+="()"
    }
    preorder(root.Left)
    preorder(root.Right)
    res+=")"
    return
}
