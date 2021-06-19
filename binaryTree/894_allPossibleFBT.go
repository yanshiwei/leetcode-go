/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    /*
    不要想那么多，纯递归思想！如果你想着怎么构建左子树右子树，就变的无从下手了！
    树类题目，不好思考就想着纯递归
    递归的时候不要想如果解子问题以及子子问题以及子子子... 问题，
    你就想着当前的这个树的节点，我要解决的话，就得拿到左子树和右子树的所有情况，
    然后根据两边的集合拼凑当前节点的结果就好了！然后加上递归的终止条件，就行了！
    递归就是这么简单！树天生就是可以拆分成子问题的题目！天生就可以用递
    */
    var res []*TreeNode
    if n==1{
        root:=&TreeNode{}
        res=append(res,root)
        return res
    }
    //左右子树可能包含1个-n-2个
    for i:=1;i<n-1;i++{
        //当左子树i个，右子树n-1-i时
        leftTreeList:=allPossibleFBT(i)
        rightTreeList:=allPossibleFBT(n-1-i)
        // 拿到当前子结果，计算当前节点的结果
        //左右合并
        for l:=0;l<len(leftTreeList);l++{
            for r:=0;r<len(rightTreeList);r++{
                root:=&TreeNode{}
                root.Left=leftTreeList[l]
                root.Right=rightTreeList[r]
                res=append(res,root)
            }
        }
    }
    return res
}
