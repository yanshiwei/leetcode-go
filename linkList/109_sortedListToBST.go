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
func sortedListToBST(head *ListNode) *TreeNode {
    /*
    参考：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/solution/you-xu-lian-biao-zhuan-huan-er-cha-sou-suo-shu-1-3/
    将给定的有序链表转换为二叉搜索树的第一步是确定根节点。由于我们需要构造出平衡的二叉树，因此比较直观的想法是让根节点左子树中的节点个数与右子树中的节点个数尽可能接近。这样一来，左右子树的高度也会非常接近，可以达到高度差绝对值不超过 1的题目要求。可以找出链表元素的中位数作为根节点的值。小于中位数的元素组成了左子树，大于中位数的元素组成了右子树，它们分别对应着有序链表中连续的一段。在这之后，我们使用分治的思想，继续递归地对左右子树进行构造，找出对应的中位数作为根节点，以此类推。
    找出链表中位数节点的方法多种多样，其中较为简单的一种是「快慢指针法」。指针到达边界（即快指针到达右端点或快指针的下一个节点是右端点）。此时，慢指针对应的元素就是中位数。在找出了中位数节点之后，我们将其作为当前根节点的元素，并递归地构造其左侧部分的链表对应的左子树，以及右侧部分的链表对应的右子树。
    */
    return buidlTree(head,nil)
}

func buidlTree(head,tail *ListNode)*TreeNode{
    if head==tail{
        return nil
    }
    var mid=getMidNode(head,tail)
    var root=&TreeNode{
        Val:mid.Val,
    }
    root.Left=buidlTree(head,mid)
    root.Right=buidlTree(mid.Next,tail)
    return root
}
//初始时，快指针fast 和慢指针 slow 均指向链表的左端点left。我们将快指针fast 向右移动两次的同时，将慢指针 slow 向右移动一次，直到快指针到达边界（即快指针到达右端点此时对应偶数长或快指针的下一个节点是右端点对于奇数长）。此时，慢指针对应的元素就是中位数
func getMidNode(head,tail *ListNode) *ListNode{
    var fast=head
    var slow=head
    for fast!=tail&&fast.Next!=tail{
        fast=fast.Next.Next
        slow=slow.Next
    }
    return slow
}
