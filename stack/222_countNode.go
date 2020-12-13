/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //递归 这是统计树节点数常用做法
func countNodes(root *TreeNode) int {
    if root==nil {
        return 0
    }
    if root.Left==nil &&root.Right==nil {
        return 1
    }
    leftCnt:=countNodes(root.Left)
    rightcnt:=countNodes(root.Right)
    return leftCnt+rightcnt+1
}
