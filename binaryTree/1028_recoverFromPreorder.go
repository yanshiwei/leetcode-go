/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var curIdx int
 var curD int
func recoverFromPreorder(traversal string) *TreeNode {
    //参考https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/solution/dian-xing-de-er-cha-shu-ti-xing-by-fresh_zhe/
    //所以我们只需要对S字符串进行全局遍历即可。同一个全局变量curIdx去处理节点值的位置，同时用一个全局变量curD维护当前遍历到的深度，再用一个栈变量depth维护当前节点深度，如果curD>depth则显然全局遍历的下一节点为当前节点的子节点，否则为其长辈节点
    curIdx=0
    curD=0
    return dfs(0,traversal)
}

func dfs(depth int,nodes string)*TreeNode{
    //fmt.Println(depth,curIdx)
    var val int
    for curIdx<len(nodes)&&nodes[curIdx]!='-'{
        v,_:=strconv.Atoi(string(nodes[curIdx]))
        val=val*10+v
        curIdx++
    }
    curD=0
    for curIdx<len(nodes)&&nodes[curIdx]=='-'{
        curIdx++
        curD++
    }
    root:=&TreeNode{Val:val}
    if curD>depth{
        fmt.Println(depth,curIdx,"left")
        //下一节点为当前节点的左子节点
        root.Left=dfs(curD,nodes)
    }
    //此时curD 由于是全局变量，有可能被上一个dfs修改，故需要重新判断curD>depth,即判断下一个节点是否是当前节点的右子节点
    if curD>depth{
        //下一节点为当前节点的子节点
        fmt.Println(depth,curIdx,"right")
        root.Right=dfs(curD,nodes)
    }
    return root    
}
