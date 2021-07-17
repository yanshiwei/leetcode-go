/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A==nil||B==nil{
        return false
    }
    return dfs(A,B)||isSubStructure(A.Left,B)||isSubStructure(A.Right,B)
}

func dfs(A *TreeNode, B *TreeNode)bool{
    if B==nil{
        //B 是空树，是任何树的子树
        return true
    }
    if A ==nil{
        //A是空树，不能说任何子树的父树
        return false
    }
    if A.Val!=B.Val{
        return false
    }
    return dfs(A.Left,B.Left)&&dfs(A.Right,B.Right)
}
