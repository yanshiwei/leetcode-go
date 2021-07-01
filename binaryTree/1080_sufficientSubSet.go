/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    if root==nil{
        return nil
    }
    if root.Left==nil&&root.Right==nil{
        //对叶子节点来说只要值的和小于limit就不存在
        if root.Val<limit{
            return nil
        }else{
            return root
        }
    }else{
        //对于分支节点来说只有当左边或右边有一个可以存在即当前分支节点就可以存在
        var leftFlag bool
        if root.Left!=nil{
            root.Left.Val+=root.Val
            leftTree:=sufficientSubset(root.Left,limit)
            root.Left.Val-=root.Val//恢复现场
            if leftTree!=nil{
                leftFlag=true
            }
            if leftFlag==false{
                //左子树不存在
                root.Left=nil
            }
        }
        var ritghFlag bool
        if root.Right!=nil{
            root.Right.Val+=root.Val
            rightTree:=sufficientSubset(root.Right,limit)
            root.Right.Val-=root.Val
            if rightTree!=nil{
                ritghFlag=true
            }
            if ritghFlag==false{
                //右子树不存在
                root.Right=nil
            }
        }
        if leftFlag||ritghFlag{
            //只有当左边或右边有一个可以存在即当前分支节点就可以存在
            return root
        }
        //左右都空，则返回空
        return nil
    }
}
