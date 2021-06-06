/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var f,g map[*TreeNode]int
func rob(root *TreeNode) int {
    /*
    一棵二叉树，树上的每个点都有对应的权值，每个点有两种状态（选中和不选中），问在不能同时选中有父子关系的点的情况下，能选中的点的最大权值和是多少：
    可以用 f(o) 表示选择 o 节点的情况下，o 节点的子树上被选择的节点的最大权值和；g(o) 表示不选择 o 节点的情况下，o 节点的子树上被选择的节点的最大权值和；l 和 r 代表 o 的左右孩子
    （1）当 o 被选中时，o 的左右孩子都不能被选中，故 o 被选中情况下子树上被选中点的最大权值和为 l 和 r 不被选中的最大权值和相加，即 f(o) = g(l) + g(r)；
    （2）当 o 不被选中时，o 的左右孩子可以被选中，也可以不被选中。对于 o 的某个具体的孩子 x，它对 o 的贡献是 x 被选中和不被选中情况下权值和的较大值。故 g(o)=max{f(l),g(l)}+max{f(r),g(r)}
    至此，我们可以用哈希表来存 f 和 g 的函数值，用深度优先搜索的办法后序遍历这棵二叉树，我们就可以得到每一个节点的 f 和 g。根节点的 f 和 g 的最大值就是我们要找的答案
    初始化为0
    */
    f=make(map[*TreeNode]int)
    g=make(map[*TreeNode]int)
    dfs(root)
    return max(f[root],g[root])
}

func dfs(root *TreeNode){
    //DFS之后序遍历
    if root==nil{
        return
    }
    dfs(root.Left)
    dfs(root.Right)
    f[root]=root.Val+g[root.Left]+g[root.Right]
    g[root]=max(f[root.Left],g[root.Left])+max(f[root.Right],g[root.Right])
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
