/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    /*
    类似 100 https://leetcode-cn.com/problems/same-tree/
    只不过次题是求子树是否包含，这个题的做法就是在 s 的每个子节点上，判断该子节点是否和 t 相等。
    判断 t 是否为 s 的子树的三个条件是或的关系，即：
    1. 当前两个树一样
    2. t是s的左子树
    3. t是s的右子树
    */
    if root==nil&&subRoot==nil{
        return true
    }
    if root==nil||subRoot==nil{
        return false
    }
    //表示任从根的何一个节点开始的子节点 而不仅仅从根
    return isSameTree(root,subRoot)||isSubtree(root.Left,subRoot)||isSubtree(root.Right,subRoot)
}

func isSameTree(p,q *TreeNode)bool{
    if p==nil&&q==nil{
        return true
    }
    if p==nil||q==nil{
        return false
    }
    if p.Val!=q.Val{
        return false
    }
    return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
}
