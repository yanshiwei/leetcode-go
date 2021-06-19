var res []int
var childCnt []int
var pathLen []int
var graph[][]int
var nodeNum int
func sumOfDistancesInTree(n int, edges [][]int) []int {
    //参考：https://leetcode-cn.com/problems/sum-of-distances-in-tree/solution/shou-hua-tu-jie-shu-zhong-ju-chi-zhi-he-shu-xing-d/
    nodeNum=n
    res=make([]int,nodeNum)
    childCnt=make([]int,nodeNum)
    pathLen=make([]int,nodeNum)
    graph=make([][]int,nodeNum)
    for i:=range edges{
        edge:=edges[i]
        graph[edge[0]]=append(graph[edge[0]],edge[1])
        graph[edge[1]]=append(graph[edge[1]],edge[0])
    }
    dfsForPath(0,-1)
    dfsParent(0,-1)
    return res
}
//第一次遍历的时候，求出每个节点为根的子树的节点个数以及其到其为根的子树中的所有节点的路径和
func dfsForPath(cur int, parent int){
    for i:=range graph[cur]{
        //对于cur的一个子节点x
        x:=graph[cur][i]
        if x==parent{
            continue
        }
        dfsForPath(x,cur)
        childCnt[cur]+=childCnt[x]//加上子树的节点个数
        pathLen[cur]+=pathLen[x]//加上子树的路径和
        pathLen[cur]+=childCnt[x]
    }
    childCnt[cur]++//加上本身节点
    return
}
//第二次遍历的时候，每个节点利用其父亲节点的信息，来求出其通过父亲节点才能到达的其它节点的路径和
func dfsParent(cur int,parent int){
    if parent==-1{
        res[cur]=pathLen[cur]
    }else{
        res[cur]=res[parent]-childCnt[cur]+nodeNum-childCnt[cur]
    }
    for i:=range graph[cur]{
        x:=graph[cur][i]
        if x==parent{
            continue
        }
        dfsParent(x,cur)
    }
}
