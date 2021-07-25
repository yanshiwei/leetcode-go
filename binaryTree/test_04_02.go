/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 /*
 因为是个有序数组，并且是要构成搜索二叉树。搜索二叉的特点就是根节点大于左节点大于右节点。既然是要构成深度最小的数，那么数就应该尽可能的饱满。所以我们就选择数组的中间点，那么左边和右边都是同样的大小
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums)<1{
        return nil
    }
    var mid=nums[len(nums)/2]
    var idx=len(nums)/2
    var root=&TreeNode{Val:mid}
    root.Left=sortedArrayToBST(nums[:idx])
    root.Right=sortedArrayToBST(nums[idx+1:])
    return root
}
