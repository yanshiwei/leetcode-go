/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    /*
    参考：https://leetcode-cn.com/problems/delete-node-in-a-bst/solution/miao-dong-jiu-wan-shi-liao-by-terry2020-tc0o/
    根据BST性质：
    1 若目标节点大于当前节点值，去右子树
    2 若目标节点小于当前节点值，去左子树
    3 若目标节点值等于当前节点的：接下来分三种情况区分：
        3.1 无左子树，当前节点的右子树直接顶替，删除当前节点
        3.2 无右子树，当前节点的左子树直接顶替，删除当前节点
        3.3 左右均有，将左子树转移到右子树的最左节点的左子树上，然后右子树顶替，删除节点
    */
    if root==nil{
        return nil
    }
    if key>root.Val{
        //去右子树
        root.Right=deleteNode(root.Right,key)
    }else if key<root.Val{
        //去左子树
        root.Left=deleteNode(root.Left,key)
    }else{
        //左右均有
        //3.1 无左子树，当前节点的右子树直接顶替，删除当前节点
        if root.Left==nil{
            return root.Right
        }
        //3.2 无右子树，当前节点的左子树直接顶替，删除当前节点
        if root.Right==nil{
            return root.Left
        }
        //3.3 左右均有，将左子树转移到右子树的最左节点的左子树上，然后右子树顶替，删除节点
        cur:=root.Right
        for cur.Left!=nil{
            cur=cur.Left
        }
        cur.Left=root.Left
        root=root.Right
    }
    return root
}
