/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var fre map[string]int
var res []*TreeNode
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    //使用深度优先搜索，其中递归函数返回当前子树的序列化结果。把每个节点开始的子树序列化结果保存在 mapmap 中，然后判断是否存在重复的子树
    fre=make(map[string]int)
    res=nil
    dfs(root)
    return res
}

func dfs(root *TreeNode)string{
    if root==nil{
        return "nil"
    }
    str:=strconv.Itoa(root.Val)
    str+=","
    leftStr:=dfs(root.Left)
    rightStr:=dfs(root.Right)
    str+=leftStr
    str+=rightStr
    fre[str]++
    if fre[str]==2{
        //第一次重复
        res=append(res,root)
    }
    return str
}
