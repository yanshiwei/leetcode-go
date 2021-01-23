/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
 原理：
    由后序遍历获取根节点；然后在中序遍历中找到根节点的位置，它的左边就是左子树的节点，右边就是右子树的节点。生成左子树和右子树就可以递归的进行
步骤：
    0 由start end指针指向前与中序列范围
    1 先由前找到根节点；
    2 遍历中，得到根节点位置idx及左子树和右子树，得到左子树数目和右子树数目，
    3.1对于左子树部分：
        后序列范围为start=start，end=start+左子树数目；
        中序列范围start=start,end=根节点位置idx
    3.2对于右子树部分：
        后序列范围为start=start+左子树数目，end=end-1;
        中序列范围start=根节点位置idx+1，end=end
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    return helper(postorder,inorder,0,len(postorder),0,len(inorder))
}

func helper(postorder[]int,inorder[]int,pStart,pend int,iStart,iEnd int)*TreeNode{
    if pStart==pend{
        //postorder 为空
        return nil
    }
    var rootV=postorder[pend-1]
    var root=&TreeNode{Val:rootV}
    //中序遍历找到根节点位置
    var idxInorder int
    for i:=iStart;i<iEnd;i++{
        if rootV==inorder[i]{
            idxInorder=i
            break
        }
    }
    var leftNum=idxInorder-iStart
    //递归构造左子树
    root.Left=helper(postorder,inorder,pStart,pStart+leftNum,iStart,idxInorder)
    //递归构造右子树
    root.Right=helper(postorder,inorder,pStart+leftNum,pend-1,idxInorder+1,iEnd)
    return root

}
