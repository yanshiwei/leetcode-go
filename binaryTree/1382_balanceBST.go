/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var inOrder []int
func balanceBST(root *TreeNode) *TreeNode {
    inOrder=nil
    //二叉搜索树的中序遍历是升序序列
    inOrderByDFS(root)
    //二分法来构造平衡的二叉搜索树,proof see https://leetcode-cn.com/problems/balance-a-binary-search-tree/solution/jiang-er-cha-sou-suo-shu-bian-ping-heng-by-leetcod/
    return buildAVL(0,len(inOrder)-1)
}
func inOrderByDFS(root *TreeNode){
    if root==nil{
        return
    }
    inOrderByDFS(root.Left)
    inOrder=append(inOrder,root.Val)
    inOrderByDFS(root.Right)
}

func buildAVL(left,right int)*TreeNode{
    var mid=left+(right-left)/2
    var node=&TreeNode{Val:inOrder[mid]}
    if left<mid{
        node.Left=buildAVL(left,mid-1)
    }
    if mid<right{
        node.Right=buildAVL(mid+1,right)
    }
    return node
}
