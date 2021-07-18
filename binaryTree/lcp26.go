/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func navigation(root *TreeNode) int {
    /*
    每个三叉点（该节点既有不为空的左右子树，又有父亲节点）的三条支路中，至少两条支路上有导航装置（
    */
    res=0
    if root==nil{
        return 0
    }
    left:=dfs(root.Left)
    right:=dfs(root.Right)
        //下面这行代码是多种状态综合处理的表现形式，具体如下：
        //左右孩子状态集(左右交换就不写了)：
        //[0,0]：左右都没放导航(包含子树为空的情况)，那么要必须放一个导航，即res+1，而且总共放一个导航；
        //[0,1]：一个孩子没放导航，再和根节点放一起，视为最近的三叉点的一条支路。由于另一个孩子状态为1，表示该三叉点有两条支路未放导航，所以还要再增加一个导航，那么res+1；
        //[0,2]：一个孩子没放导航，再和根节点放一起，视为最近的三叉点的一条支路。另一个孩子状态为2，表示该三叉点有两条支路已放导航，所以不用再增加导航，返回res即可；
        //[1,1]：其中一个孩子和根节点放一起，视为一条支路，那么就表示该合并后的支路上有导航（只要有就可以了），所以该三叉点的两条支路有导航，返回res即可；
        //[1,2]：前面同上，该三叉点的三条支路都有导航，返回res即可；
        //[2,2]：直接同上。
        //综合上面分类情况，左右返回的状态值相加left+right>=2时，返回res即可，其他情况都要再加一个导航。
        if left+right>=2{
            return res
        }
        return res+1
}
// 0 代表该点及孩子节点都不会放导航装置
// 1 代表左右孩子都没有导航装置时，必须选一个孩子支路放一个导航装置，另一个暂时不放，以1状态标记交上面节点去处理
// 2 表示左右孩子支路都有导航装置，该节点的父节点支路可以放也可以不放导航装置
func dfs(root *TreeNode)int{
    if root==nil{
        return 0
    }
    left:=dfs(root.Left)
    right:=dfs(root.Right)
    //这里只要左右不为空，即为三叉点
    if root.Left!=nil&&root.Right!=nil{
         //左右子树都没放导航，那么必须选一条支路放一个导航，另一条支路暂时不放，返回状态1
         if left==0&&right==0{
             res+=1
             return 1
         }
         //一条支路有导航，另一条支路没有导航，继续暂时不放，返回状态1，要不要加交给上面节点来判断处理；
         if left==0||right==0{
             return 1
         }
         //左右孩子支路都有导航，那么就返回状态2
         return 2
    }else if root.Left==nil{
        //左孩子为空，该节点状态等于右孩子
        return right
    }else if root.Right==nil{
        //like
        return left
    }
    return 0
}


