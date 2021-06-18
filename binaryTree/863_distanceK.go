/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res[]int
 var fre map[*TreeNode]*TreeNode//// 父节点的映射表
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    /*
    思路分析

    1.按照例子给出的结果，这里遍历不是简单子节点，也可能是父节点的子节点（如例子里的1）
    2.为了查找父节点，需要构建一个额外的map来保存节点的父节点，这样子就可以做三个方向的遍历：父节点，左子节点和右子节点
    3.从target节点从三个方向做广度优先的遍历，每次距离加1，正好为k的时候，在queue里的节点就是答案
    4.坑： 这里遍历会三个方向存在循环的问题，所以需要记录哪些结点已经遍历了
    */
    res=nil
    fre=make(map[*TreeNode]*TreeNode)
    // 构建好映射表
    dfs(root,nil)
    //bfs 
    var queue []*TreeNode
    // 避免死循环的记录遍历的点
    var visited=make(map[*TreeNode]bool)
    //从target开始
    queue=append(queue,target)
    visited[target]=true
    var curLevel int// 记录当前遍历的层数，其实就是距离
    for len(queue)>0{
        if curLevel==k{
            //该层所有元素与target距离为k
            for len(queue)>0{
                head:=queue[0]
                res=append(res,head.Val)
                queue=queue[1:]
            }
            return res
        }
        curLen:=len(queue)
        for i:=0;i<curLen;i++{
            head:=queue[0]
            queue=queue[1:]
            //fmt.Println(head.Val)
            // 往三个方向去遍历
            if head.Left!=nil{
                if _,ok:=visited[head.Left];ok==false{
                    queue=append(queue,head.Left)
                    visited[head.Left]=true
                }
            }
            if head.Right!=nil{
                if _,ok:=visited[head.Right];ok==false{
                    queue=append(queue,head.Right)
                    visited[head.Right]=true
                }
            }
            pa:=fre[head]
            if pa!=nil{
                if _,ok:=visited[pa];ok==false{
                    queue=append(queue,pa)
                    visited[pa]=true
                }
            }           
        }
        // 一层遍历完成距离+1
        curLevel++
    }
    return res
}
// 递归方式来构建父节点信息
func dfs(cur *TreeNode, parent *TreeNode){
    if cur==nil{
        return
    }
    fre[cur]=parent
    dfs(cur.Left,cur)
    dfs(cur.Right,cur)
}
