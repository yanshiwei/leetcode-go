func findRedundantDirectedConnection(edges [][]int) []int {
    /*参考https://leetcode-cn.com/problems/redundant-connection-ii/solution/bing-cha-ji-jie-fa-bu-xu-yao-bao-cun-zhe-1uom/
    树中的每个节点都有一个父节点，除了根节点没有父节点。在多了一条附加的边之后，可能有以下两种情况：
    1. 附加的边指向根节点，则包括根节点在内的每个节点都有一个父节点，此时图中一定有环路；
    2. 附加的边指向非根节点，则恰好有一个节点（即被附加的边指向的节点）有两个父节点，此时图中可能有环路也可能没有环路.此时图中可能有环路也可能没有环路!
    对应策略：
    1. 由于只存在一条附加边，对于第一种情况，所有点的入度都为1，此时，我们只需要删除环上的任一条边，均可形成‘有根树’，根据题目要求，我们选择数组中最后一条边；
    2. 对于第二种情况，添加附加边后，该附加边指向的节点入度为2，且根结点入度为1，其他节点入度为1，因此我们只需要找到入度为2的节点，其中一条指向该点的边即为附加边。
    流程：
    1. 通过所有节点的入度判断属于那种情况，如果存在一个点的入度为2，则属于情形2，并保存到以该点为弧头的两条边（答案一定存在于其中一条），否则属于情形1.
    2. 创建并查集，用来检查连通性
    3. 如果属于情形1，数组中依次添加边[u,v]，如果father(u)=v,则直接输出结果[u,v]
    4. 如果属于情形2，保存以该入度为2为弧头的两条边，将数组中其余n-2条边合并至并查集，此时的图一定是无环图，最后依次添加这两条边，其中成环的边即为附加边，也就是情形2的结果
    */
    nodeCnt:=len(edges)
    //并查集-初始化
    var fa=make([]int,nodeCnt+1)
    for i:=1;i<=nodeCnt;i++{
        fa[i]=i
    }
    var degree=make([]int,nodeCnt+1)
    var flag=-1//默认节点的度
    for i:=range edges{
        edge:=edges[i]
        //node1:=edge[0]
        node2:=edge[1]
        degree[node2]++
        if degree[node2]==2{
            flag=node2
            break
        }
    }
    //保留情况2的两条边
    var cacheEdges [][]int
    for i:=range edges{
        edge:=edges[i]
        node1:=edge[0]
        node2:=edge[1]
        if flag==node2{
            //不等于-1时，属于情形2，并保存以flag为入度的边
            cacheEdges=append(cacheEdges,edge)
        }else{
            //返回情形1的结果，情况1一定有环，直接并查集 并接口，如有环，则并之前已经是相同一个集合
            if unoin(fa,node1,node2)==false{
                //有环
                return []int{node1,node2}
            }
        }
    }
    //返回情形2的结果
    for i:=range cacheEdges{
        edge:=cacheEdges[i]
        node1:=edge[0]
        node2:=edge[1]
        if unoin(fa,node1,node2)==false{
            //有环
            return []int{node1,node2}
        }else{
            //删除任意一个，根据题目要求，我们选择数组中最后一条边
            return cacheEdges[1]
        }
    }
    return []int{}
}

//并查集-并
func unoin(fa[]int,i,j int)bool{
    x:=find(fa,i)
    y:=find(fa,j)
    if x!=y{
        fa[y]=x
        return true
    }else{
        return false
    }
}
//并查集-查
func find(fa[]int,x int)int {
    if fa[x]==x{
        return x
    }else{
        return find(fa,fa[x])
    }
}
