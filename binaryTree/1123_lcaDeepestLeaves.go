/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    //深度相同，说明最深的节点在这个节点两边，那这个节点就是结果,这是重点
    if root==nil{
        return root
    }
    leftV:=getMaxDepth(root.Left)
    rightV:=getMaxDepth(root.Right)
    if leftV==rightV{
        return root
    }
    if leftV>rightV{
        return lcaDeepestLeaves(root.Left)
    }
    return lcaDeepestLeaves(root.Right)
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
