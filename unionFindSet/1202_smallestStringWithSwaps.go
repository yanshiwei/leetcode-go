func smallestStringWithSwaps(s string, pairs [][]int) string {
    /*
    并查集，参考：https://leetcode-cn.com/problems/smallest-string-with-swaps/solution/bing-cha-ji-sheng-cheng-lian-tong-tu-by-ea8er/
    思路：
    将每一个字符抽象为「点」，那么这些「索引对」即为「边」，我们只需要维护这个「图」的连通性即可。对于同属一个连通块（极大连通子图）内的字符，我们可以任意地交换它们。
    这样我们的思路就很清晰了：利用并查集维护任意两点的连通性，将同属一个连通块内的点提取出来，直接排序后放置回其在字符串中的原位置即可。
    具体做法：
    1.使用并查集，构造图
    2.遍历所有非连通图
    3.遍历parent数组，建立一个map来保存连通图，key为根节点值，value为数组，存储以该根节点为根的并查集中所有节点下标
    4.对map中每个数组(同一连通图)元素按字典序排序
    5.将排序好的元素依次插入到结果中
    注意并查集操作的是下标！
    */
    var n=len(s)
    var ds=NewUFSet(n)
    var rs=make([]byte,n)
    for _,v:=range pairs{
        ds.Merge(v[0],v[1])
    }
    record:=make(map[int][]int)// key is root, value is 以该ROOT的所有节点
    for i:=0;i<n;i++{
        record[ds.Find(i)]=append(record[ds.Find(i)],i)
    }
    for _,v :=range record{
        c:=make([]int,len(v))
        copy(c,v)
        //元素按字典序排序
        sort.Slice(v,func(i,j int)bool{
            return s[v[i]]<s[v[j]]
        })
        for i := 0; i < len(c); i++ {
            rs[c[i]] = s[v[i]]
        }
    }
    return string(rs)
}

//并查集实现
type UFSet struct {
    Father []int
    Rank []int
}

func NewUFSet(n int)UFSet{
    father:=make([]int,n)
    rank:=make([]int,n)
    for i:=0;i<n;i++{
        father[i]=i
        rank[i]=0
    }
    return UFSet{
        Father:father,
        Rank:rank,
    }
}
// 路径压缩， 遍历过程中的所有父节点直接指向根节点，
// 减少后续查找次数
func (ds UFSet)Find(x int)int {
    if x==ds.Father[x]{
        return x
    }
    ds.Father[x]=ds.Find(ds.Father[x])//父节点设为根节点，路径压缩
    return ds.Father[x]
}
//合并两个节点
// 如果处于同一个并查集， 不需要合并
// 如果不处于同一个并查集，判断两个rootx和rooty谁的秩大
func (ds UFSet)Merge(x, y int){
    rootX:=ds.Find(x)
    rootY:=ds.Find(y)
    if rootX!=rootY{
        //按rank合并，合并到大的rank
        if ds.Rank[rootX]<ds.Rank[rootY]{
            rootX,rootY=rootY,rootX
        }
        ds.Father[rootY]=rootX
        if ds.Rank[rootX]==ds.Rank[rootY]{
            ds.Rank[rootX]++
        }
    }
}
