/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res []int
 var idx int
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    res=nil
    idx=0
    if dfs(root,voyage)==true{
        return res
    }
    return []int{-1}
}
func dfs(root *TreeNode,voyage []int)bool{
    if root==nil{
        //如果树为空，肯定可以匹配
        return true
    }
    //值不等，肯定无法通过翻转来达到匹配，或者数组索引越界结束递归
    if root.Val!=voyage[idx]||idx+1>len(voyage){
        return false
    }
    idx++
    //如果匹配，继续递归尝试匹配左子树和右子树
    if dfs(root.Left,voyage)&&dfs(root.Right,voyage){
        //如果同时匹配返回匹配成功
        return true
    }
    //如果存在子树不匹配，尝试翻转左右子树，注意这里不是真翻转，而是对比时交换左右顺序即可,voyage顺序不变
    if dfs(root.Right,voyage)&&dfs(root.Left,voyage){
        //通过翻转能够匹配，保存当前root->val到结果res
        res=append(res,root.Val)
        return true
    }
    //翻也无法匹配，则最终匹配失败
    return false
}
