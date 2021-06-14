/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var forest []*TreeNode// 存储最终的森林
 var fre map[int]bool// 存储待删除节点值
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    forest=nil
    fre=make(map[int]bool)
    // 将待删除的节点的值存到map中，便于直接使用判断
    for i:=range to_delete{
        fre[to_delete[i]]=true
    }
    // 返回删除后的新二叉树，若不为空，则也添加到森林中
    root=dfs(root)
    if root!=nil{
        forest=append(forest,root)
    }
    return forest
}
/*
递归三要素：
    1,递归结束条件：二叉树已空
    2,本级递归做的事情：如果当前节点存在于待删除列表中，则将其非空子树存储到forest中，并移除该节点
    3,本级递归返回值：移除掉特定节点后的二叉树
*/
//post order
func dfs(root *TreeNode)*TreeNode{
    if root==nil{
        return nil
    }
    // 去左、右子树中移除`to_delete`中出现的节点
    root.Left=dfs(root.Left)
    root.Right=dfs(root.Right)
    // 本级递归做的事：如果当前节点存在于待删除列表中，则将其非空子树存储到forest中，并移除该节点
    if _,ok:=fre[root.Val];ok{
        if root.Left!=nil{
            forest=append(forest,root.Left)
        }
        if root.Right!=nil{
            forest=append(forest,root.Right)
        }
        root=nil//删除
    }
    return root
}
