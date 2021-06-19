/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    /*
    从每个树开始，获得当前节点的左右子树的最大深度,
    1.深度相同，说明最深的节点在这个节点两边，那这个节点就是结果,这是重点
    2.如果深度不相同，则去深度大的子树继续判断，最终就能得到结果
    */
    if root==nil{
        return root
    }
    // 获取当前节点的左右子树的最大深度
    leftDepth:=getMaxDepth(root.Left)
    rightDepth:=getMaxDepth(root.Right)
     // 如果两边最大深度相同，则这个节点就是结果
     if leftDepth==rightDepth{
         return root
     }
     // 不相等，那就去深度大的子树那边继续找
     if leftDepth>rightDepth{
         return subtreeWithAllDeepest(root.Left)
     }
     return subtreeWithAllDeepest(root.Right)
}

func getMaxDepth(root *TreeNode)int {
    if root==nil{
        return 0
    }
    return max(getMaxDepth(root.Left),getMaxDepth(root.Right))+1
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
