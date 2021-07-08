func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    var visited=make([]bool,n)// 统计出每个结点 i 的入度 
    var ds=NewUFSet(n)
    var count=n
    for i:=0;i<n;i++{
        if leftChild[i]!=-1{
            //可构造节点
            if visited[leftChild[i]]{
                //判断是否只有一个前驱
                return false
            }
            if ds.Merge(i,leftChild[i]){
                //同一个集合,已经存在于一个集合中，那么说明会产生环，返回错误
                return false
            }
            visited[leftChild[i]]=true
            count--
        }
        if rightChild[i]!=-1{
            if visited[rightChild[i]]{
                return false
            }
            if ds.Merge(i,rightChild[i]){
                return false
            }
            visited[rightChild[i]]=true
            count--
        }
    }
    //fmt.Print(ds)
    //如果最后所有的节点组成一个连通分支，才是一棵树
    if ds.Cnt!=1{
        return false
    }
    return count == 1;//判断是否只有一个根节点（即前驱数量为零的节点）like '3[1,-1,-1] [-1,-1,1]'
}

type UFSet struct {
    Father []int
    Rank []int
    Cnt int
}
func NewUFSet(n int)*UFSet{
    father:=make([]int,n)
    rank:=make([]int,n)
    for i:=0;i<n;i++{
        father[i]=i
        rank[i]=1
    }
    return &UFSet{
        Father:father,
        Rank:rank,
        Cnt:n,
    }
}
//路径压缩， 遍历过程中的所有父节点直接指向根节点，
// 减少后续查找次数
func (ds* UFSet)Find(x int)int {
    if x==ds.Father[x]{
        return x
    }
    ds.Father[x]=ds.Find(ds.Father[x])//父节点设为根节点，路径压缩
    return ds.Father[x]
}
//合并两个节点
// 如果处于同一个并查集， 不需要合并
// 如果不处于同一个并查集，判断两个rootx和rooty谁的秩大
func (ds* UFSet)Merge(x,y int)bool{
    rootX:=ds.Find(x)
    rootY:=ds.Find(y)
    if rootX!=rootY{
         //按rank合并，合并到大的rank
         if ds.Rank[rootX]<ds.Rank[rootY]{
             ds.Father[rootX]=rootY
         }else{
             ds.Father[rootY]=rootX
         }
         if ds.Rank[rootX]==ds.Rank[rootY]{
             ds.Rank[rootX]++ //如果深度相同且根节点不同，则新的根节点的深度+1
         }
         ds.Cnt--
         return false
    }
    return true
}
