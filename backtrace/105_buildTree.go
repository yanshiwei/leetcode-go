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
    由前遍历获取根节点；然后在中序遍历中找到根节点的位置，它的左边就是左子树的节点，右边就是右子树的节点。生成左子树和右子树就可以递归的进行
步骤：
    0 由start end指针指向前与中序列范围
    1 先由前找到根节点；
    2 遍历中，得到根节点位置idx及左子树和右子树，得到左子树数目和右子树数目，
    3.1对于左子树部分：
        前序列范围为start=start+1（之所以+1是过掉第一个根节点），end=start+1+左子树数目；
        中序列范围start=start,end=根节点位置idx
    3.2对于右子树部分：
        前序列范围为start=start+1+左子树数目，end=end;
        中序列范围start=根节点位置idx+1，end=end
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    return helper(preorder,inorder,0,len(preorder),0,len(inorder))
}

func helper(preorder[]int,inorder[]int,pStart,pend int,iStart,iEnd int)*TreeNode{
    if pStart==pend{
        //preorder 为空
        return nil
    }
    var rootV=preorder[pStart]
    var root=&TreeNode{Val:rootV}
    //preorder遍历找到根节点位置
    var idxInorder int
    for i:=iStart;i<iEnd;i++{
        if rootV==inorder[i]{
            idxInorder=i
            break
        }
    }
    var leftNum=idxInorder-iStart
    //递归构造左子树
    root.Left=helper(preorder,inorder,pStart+1,pStart+1+leftNum,iStart,idxInorder)
    //递归构造右子树
    root.Right=helper(preorder,inorder,pStart+1+leftNum,pend,idxInorder+1,iEnd)
    return root

}
