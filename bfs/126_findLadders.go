func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    var wordId=make(map[string]int)// 单词到id的映射
    var idWord []string// id到单词的映射
    var res [][]string
    var edges[][]int// 图的边
     // 将wordList所有单词加入wordId中 相同的只保留一个 
     // 并为每一个单词分配一个id
    var i int
    for i=range wordList {
        if _,ok:=wordId[wordList[i]];ok==false {
            wordId[wordList[i]]=i
            idWord=append(idWord,wordList[i])
        }
    }
    // 若endWord不在wordList中 则无解
    if _,ok:=wordId[endWord];ok==false {
        return res
    }
    i++
    // 把beginWord也加入wordId中
    if _,ok:=wordId[beginWord];ok==false{
        wordId[beginWord]=i
        idWord=append(idWord,beginWord)
    }
    //初始化边
    edges=make([][]int,len(idWord))
    for i:=0;i<len(idWord);i++ {
        for j:=i+1;j<len(idWord);j++ {
            if isNear(idWord[i],idWord[j]) {
                edges[i]=append(edges[i],j)
                edges[j]=append(edges[j],i)
            }
        }
    }
    var destId=wordId[endWord]// destId in edge
    var queue [][]int//队列的每一个元素都是一个数组，每个数组代表中保存从起点开始的所有路径，数组中的最后一个元素为当前路径的最后节点 last
    var cost=make([]int,len(idWord))//cost 数组，其中 cost[i] 表示 beginWord 对应的点到第 i 个点的代价（即转换次数）
    for i:=range cost {
        cost[i]=INTMAX
    }
    cost[wordId[beginWord]]=0//beginWord到本身为0
    queue=append(queue,[]int{wordId[beginWord]})
    for ii:=0;ii<len(queue);ii++ {
        nowSlice:=queue[ii]
        last:=nowSlice[len(nowSlice)-1]
        if last==destId {
            //该节点为终点
            var tmp []string
            for i:=range nowSlice {
                tmp=append(tmp,idWord[nowSlice[i]])
            }
            res=append(res,tmp)
        }else {
            //该节点不为终点，则遍历和它连通的节点（假设为 to ）
            for i:=0;i<len(edges[last]);i++ {
                to:=edges[last][i]
                if cost[last]+1<=cost[to]{
                    cost[to]=cost[last]+1
                    var tmp =make([]int,len(nowSlice))
                    copy(tmp,nowSlice)
                    tmp=append(tmp,to)
                    queue=append(queue,tmp)
                }
            }
        }
    }
    return res
}
//转化判断
func isNear(s,t string)bool {
    var m=len(s)
    var n=len(t)
    if m!=n {
        return false
    }
    var diff int
    for i:=0;i<m;i++ {
        if s[i]!=t[i] {
            diff++
        }
    }
    return diff==1
}

const INTMAX=int(^uint(0)>>1)
