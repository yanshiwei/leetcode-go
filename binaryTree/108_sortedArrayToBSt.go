/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    /*
    BST的中序遍历是升序的，因此本题等同于根据中序遍历的序列恢复二叉搜索树。因此我们可以以升序序列中的任一个元素作为根节点，以该元素左边的升序序列构建左子树，以该元素右边的升序序列构建右子树，这样得到的树就是一棵二叉搜索树啦～ 又因为本题要求高度平衡，因此我们需要选择升序序列的中间元素作为根节点
    */
    return dfs(nums,0,len(nums)-1)
}

func dfs(nums[]int, left,right int)*TreeNode{
    if left>right{
        return nil
    }
    // 以升序数组的中间元素作为根节点 root
    var mid=left+(right-left)/2
    root:=&TreeNode{
        Val:nums[mid],
    }
    // 递归的构建 root 的左子树与右子树。
    root.Left=dfs(nums,left,mid-1)
    root.Right=dfs(nums,mid+1,right)
    return root
}
