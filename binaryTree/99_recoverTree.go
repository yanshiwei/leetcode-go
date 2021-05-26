/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var firstNode *TreeNode 
var secondNode *TreeNode
var pre=&TreeNode{Val:INTMIN}
func recoverTree(root *TreeNode)  {
    /*
    这里我们二叉树搜索树的中序遍历(中序遍历遍历元素是递增的)
    如中序遍历顺序是 4,2,3,1，我们只要找到节点 4 和节点 1 交换顺序即可！
    这里我们有个规律发现这两个节点：
    1.第一个节点，是第一个按照中序遍历时候前一个节点大于后一个节点，我们选取前一个节点，这里指节点 4；
    2.第二个节点，是在第一个节点找到之后，后面出现前一个节点大于后一个节点，我们选择后一个节点，这里指节点 1
    */
    firstNode=nil
    secondNode=nil
    pre=&TreeNode{Val:INTMIN}
    inOrder(root)
    //swap
    fmt.Println(INTMIN,firstNode.Val,secondNode.Val)
    var tmp int =firstNode.Val
    firstNode.Val=secondNode.Val
    secondNode.Val=tmp
    return
}
const INTMAX=int(^uint(0) >> 1)
const INTMIN=^INTMAX

func inOrder(root *TreeNode){
    if root==nil{
        return
    }
    inOrder(root.Left)
    if firstNode==nil&&pre.Val>root.Val{
        //第一个节点，是第一个按照中序遍历时候前一个节点大于后一个节点，我们选取前一个节点，因为按照中序遍历，前一个应该小于后一个；
        firstNode=pre
    }
    if firstNode!=nil&&pre.Val>root.Val{
        //第二个节点，是在第一个节点找到之后，后面出现前一个节点大于后一个节点，我们选择后一个节点。因为按照中序遍历，前一个应该小于后一个
        secondNode=root
    }
    pre=root
    inOrder(root.Right)
}
