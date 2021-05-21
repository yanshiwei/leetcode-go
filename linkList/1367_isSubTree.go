/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//函数负责判断以当前节点为根的子树是否符合题目要求
 func isSubBydfs(treeNode *TreeNode, head *ListNode)bool{
     if head==nil{
         // 链表已经全部匹配完，匹配成功
         return true
     }
     // 二叉树访问到了空节点，匹配失败
     if treeNode==nil{
         return false
     }
     // 当前匹配的二叉树上节点的值与链表节点的值不相等，匹配失败
     if treeNode.Val!=head.Val{
         return false
     }
     //当前节点匹配时，需要观察后续节点情况
     return isSubBydfs(treeNode.Left,head.Next)||isSubBydfs(treeNode.Right,head.Next)
 }
 //用来遍历整个树的所有结点
func isSubPath(head *ListNode, root *TreeNode) bool {
    if head==nil{
        //空链表也匹配
        return true
    }
    if root==nil{
        //空树,没有可匹配的
        return false
    }
    return isSubBydfs(root,head)||isSubPath(head,root.Left)||isSubPath(head,root.Right)
}
