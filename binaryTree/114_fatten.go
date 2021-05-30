/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    /*
    当于把一A结点的左孩子L整个完整的树替换到此结点的右孩子R上(同时把此结点原来的右孩子R接到此节点左孩子L的最右结点的右孩子上)，然后将当前结点替换为此结点的右孩子(也就是原来的L)进行迭代(注意每次要将结点的左孩子设置为nullptr)
    */
    for root!=nil{
        left:=root.Left
        if left!=nil{
            //找到左子树的最右节点
            for left!=nil&&left.Right!=nil{
                left=left.Right
            }
            left.Right=root.Right
            root.Right=root.Left
            root.Left=nil
        }
        //前结点替换为此结点的右孩子(也就是原来的L)进行迭代
        root=root.Right
    }
}
