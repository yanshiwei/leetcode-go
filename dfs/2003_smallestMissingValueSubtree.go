var length=100010
var e [][]int
var ans []int
var now int
var has1 []bool
var vis []bool
func smallestMissingValueSubtree(parents []int, nums []int) []int {
    e=make([][]int,length)//element in e is node
    has1=make([]bool,length)//element in has1 is node
    vis=make([]bool,length)//element in vis is number
    var n=len(parents)
    ans=make([]int,n)//element in ans is node
    now=1
    // 第一遍 dfs，先把哪些子树里有 1 找出来，没有1的子树最小值就是1
    for i:=1;i<n;i++{
        e[parents[i]]=append(e[parents[i]],i)
    }
    dfs1(0,nums)
    // 第二遍 dfs，计算有 1 的子树的答案
    dfs2(0,nums)
    return ans
}

func dfs1(son int,nums []int){
    if nums[son]==1{
        has1[son]=true
    }else{
        has1[son]=false
    }
    for _,fn:=range e[son]{
        dfs1(fn,nums)//fn为根节点
        if has1[son]||has1[fn]{
            //son或者其子树有1，则son就有1
            has1[son]=true
        }else{
            has1[son]=false
        }
    }
    if has1[son]==false{
        //son及其子树没有1，则最小就是1
        ans[son]=1
    }
}

func dfs2(son int,nums[]int){
    //只计算有1的子树
    if has1[son]==false{
        return
    }
    for _,fn:=range e[son]{
        if has1[fn]{
            dfs2(fn,nums)
        }
    }
    // 记录没有 1 的分支的数number，历浅节点'u'的其他子树w，利用vis数组标记这些节点的基因值，因为深节点v的缺失最小基因值可能被'u'或者'w'子树的基因值覆盖
    for _,fn:=range e[son]{
        if has1[fn]==false{
            dfs3(fn,nums)
        }
    }
    vis[nums[son]]=true
    // 计算本节点的答案
    for vis[now]==true{
        now++
    }
    ans[son]=now
}

func dfs3(son int,nums[]int){
    for _,fn:=range e[son]{
        dfs3(fn,nums)
    }
    //完成遍历的子树中出现过哪些数number
    vis[nums[son]]=true
}
