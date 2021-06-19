/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
    /*
    参考：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/solution/gen-ju-qian-xu-he-hou-xu-bian-li-gou-zao-er-cha-sh/
    */
    var n=len(pre)
    if n==0{
        return nil
    }
    root:=&TreeNode{Val:pre[0]}
    if n==1{
        return root
    }
    var leftNum int//左分支节点数目
    for i:=0;i<n;i++{
        //令左分支有 L个节点。我们知道左分支的头节点为 pre[1]，但它也出现在左分支的后序表示的最后。所以 pre[1] = post[L-1]，利用唯一性，找到相等的，就得到L真实的值
        if pre[1]==post[i]{
            leftNum=i+1
        }
    }
    //递归
    //左分支由 pre[1 : L+1] 和 post[0 : L] 重新分支
    root.Left=constructFromPrePost(pre[1:leftNum+1],post[0:leftNum])
    //右分支将由 pre[L+1 : N] 和 post[L : N-1] 重新分支
    root.Right=constructFromPrePost(pre[leftNum+1:],post[leftNum:n-1])
    return root
}
