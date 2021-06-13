/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    return dfs(nums,0,len(nums))
}
//dfs preorder
func dfs(nums[]int,l,r int)*TreeNode{
    if l==r{
        return nil
    }
    maxIdx:=findMax(nums,l,r)
    root:=&TreeNode{Val:nums[maxIdx]}
    root.Left=dfs(nums,l,maxIdx)
    root.Right=dfs(nums,maxIdx+1,r)
    return root
}

func findMax(nums []int,l,r int)int{
    var curMAxIdx=l
    for i:=l;i<r;i++{
        if nums[curMAxIdx]<nums[i]{
            curMAxIdx=i
        }
    }
    return curMAxIdx
}
