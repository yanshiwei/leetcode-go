/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    if root==nil{
        return 0
    }
    //先调用dfs函数从root开始查找路径，再调用pathsum函数到root左右子树开始查找这样实现从任意节点开始
    res:=helper(root,targetSum)
    res+=pathSum(root.Left,targetSum)
    res+=pathSum(root.Right,targetSum)
    return res
}

func helper(root *TreeNode, targetSum int)int{
    if root==nil{
        return 0
    }
    innerRes:=0
    targetSum-=root.Val
    //如果找到了一个路径全局变量就+1
    if targetSum==0{
        //fmt.Println("s",res)
        //注意不要return,因为不要求到叶节点结束,所以一条路径下面还可能有另一条
        innerRes++
    }
    innerRes+=helper(root.Left,targetSum)
    innerRes+=helper(root.Right,targetSum)
    return innerRes
}
