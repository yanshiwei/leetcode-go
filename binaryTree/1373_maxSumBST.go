/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func maxSumBST(root *TreeNode) int {
    res=0
    //https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/solution/python3-zi-di-xiang-shang-di-gui-by-yim-lwn0e/
    dfs(root)
    return res
}
// 当前子树的最小值、最大值和键值和。左子树只看最大，右子树只看最小
func dfs(root *TreeNode)(int,int,int){
    if root==nil{
        //nil is bst too, return min,max,0
        return INTMAX,INTMIN,0
    }
    lmin,lmax,lsum:=dfs(root.Left)
    rmin,rmax,rsum:=dfs(root.Right)
    if lmax<root.Val&&root.Val<rmin{
        //能够构成BST
        res=max(res,root.Val+lsum+rsum)
        return min(lmin,root.Val),max(root.Val,rmax),root.Val+lsum+rsum//考虑子树为空情况
    }
    //不能构成BST
    return INTMIN,INTMAX,0
}
func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
