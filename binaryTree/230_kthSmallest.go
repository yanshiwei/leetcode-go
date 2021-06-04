/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var arr []int
func kthSmallest(root *TreeNode, k int) int {
    //bst 中序遍历是升序
    arr=make([]int,0)
    nums:=inorder(root,arr)
    return nums[k-1]
}

func inorder(root *TreeNode, arr []int)[]int{
    if root==nil{
        return arr
    }
    arr=inorder(root.Left,arr)
    arr=append(arr,root.Val)
    arr=inorder(root.Right,arr)
    return arr
}
