/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func countPairs(root *TreeNode, distance int) int {
    res=0
    dfs(root,distance)
    return res
}
//返回node节点的所有叶子结点到node的合法距离，合法指的是距离不大于distance
//返回叶子节点距离组成的vector，每个元素表示一个叶子节点到当前节点的距离
func dfs(root *TreeNode,distance int)[]int{
    if root==nil{
        //无叶子 返回空，后面不处理
        return nil
    }
    if root.Left==nil&&root.Right==nil{
        //leaf
        return []int{1}
    }
    var ret []int
    left:=dfs(root.Left,distance)
    //更新左叶子节点到当前节点的距离：left是当前节点的左节点root.Left 和 root的叶子节点间的距离，当前节点比root.Left高一级，所以每个距离+1
    for i:=range left{
        if left[i]+1>distance{
            continue
        }
        ret=append(ret,left[i]+1)
    }
    right:=dfs(root.Right,distance)
    //更新右叶子节点到当前节点的距离
    for i:=range right{
        if right[i]+1>distance{
            continue
        }
        ret=append(ret,right[i]+1)
    }
    ////对每个节点进行配对，判断是否满足
    for l:=range left{
        for r:=range right{
            if left[l]+right[r]<=distance{
                //fmt.Println(left[l],right[r],root.Val)
                res++
            }
        }
    }
    return ret
}
