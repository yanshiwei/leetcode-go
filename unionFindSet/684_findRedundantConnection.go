func findRedundantConnection(edges [][]int) []int {
    //利用并查集来判环。即如果在两个结点合并前，这两个结点已经属于一个集合，那么这个图中就存在环路
    var nodeCnt=len(edges)
    //并查集-初始化
    var fa=make([]int,nodeCnt+1)
    for i:=1;i<=nodeCnt;i++{
        fa[i]=i
    }
    for i:=0;i<nodeCnt;i++{
        edge:=edges[i]
        node1:=edge[0]
        node2:=edge[1]
        if find(fa,node1)!=find(fa,node2){
            //两个节点不在同一个集合
            unoin(fa,node1,node2)
        }else{
            //在同一个集合，要删除
            return edge
        }
    }
    return []int{}
}

//并查集-并
func unoin(fa[]int,i,j int){
    fa[find(fa,i)]=find(fa,j)
}

//并查集-查
func find(fa[]int,x int)int {
    if fa[x]==x{
        //找到根节点
        return x
    }else{
        return find(fa,fa[x])
    }
}
