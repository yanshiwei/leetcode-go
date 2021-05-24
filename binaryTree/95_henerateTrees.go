/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    /*
    递归
    用一下查找二叉树的性质。左子树的所有值小于根节点，右子树的所有值大于根节点。
    所以如果求 1...n 的所有可能。我们只需要：
    1.把 1 作为根节点，[ ] 空作为左子树，[ 2 ... n ] 的所有可能作为右子树。
    2.2 作为根节点，[ 1 ] 作为左子树，[ 3...n ] 的所有可能作为右子树。
    3.3 作为根节点，[ 1 2 ] 的所有可能作为左子树，[ 4 ... n ] 的所有可能作为右子树，然后左子树和右子树两两组合。
    4.4 作为根节点，[ 1 2 3 ] 的所有可能作为左子树，[ 5 ... n ] 的所有可能作为右子树，然后左子树和右子树两两组合。
    ...
    n 作为根节点，[ 1... n ] 的所有可能作为左子树，[ ] 作为右子树。
    */
    if n==0{
        return nil
    }
    return getResIteration(1,n)
}

func getResIteration(start,end int)[]*TreeNode{
    var res []*TreeNode
    if start>end{
        //此时没有数字，将 null 加入结果中
        res=append(res,nil)
        return res
    }
    if start==end{
        //只有一个数字，当前数字作为一棵树加入结果中
        tmp:=&TreeNode{
            Val:start,
        }
        res=append(res,tmp)
        return res
    }
    //尝试每个数字作为根节点，i节点为根节点时，左子树i-1个节点，右子树n-i个节点
    for i:=start;i<=end;i++{
        //得到所有可能的左子树
        leftTreeList:=getResIteration(start,i-1)
        //得到所有可能的右子树
        rightTreeList:=getResIteration(i+1,end)
        //左子树右子树两两组合
        for _,left:=range leftTreeList{
            for _,right:=range rightTreeList{
                root:=&TreeNode{
                    Val:i,
                }
                root.Left=left
                root.Right=right
                res=append(res,root)
            }
        } 
    }
    return res
}
