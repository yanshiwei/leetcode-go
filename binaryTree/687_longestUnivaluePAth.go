/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func longestUnivaluePath(root *TreeNode) int {
    res=0
    if root==nil{
        return res
    }
    longestPath(root)
    return res
}
func longestPath(root *TreeNode)int {
    if root==nil{
        return 0
    }
    leftV:=longestPath(root.Left)
    rightV:=longestPath(root.Right)
    // 如果存在左子节点和根节点同值，更新左最长路径;否则左最长路径为0
    if root.Left!=nil&&root.Val==root.Left.Val{
        leftV++
    }else{
        leftV=0
    }
    // 如果存在右子节点和根节点同值，更新右最长路径;否则右最长路径为0
    if root.Right!=nil&&root.Val==root.Right.Val{
        rightV++
    }else{
        rightV=0
    }
    res=max(res,leftV+rightV)
    return max(leftV,rightV)
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
