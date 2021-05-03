func eatenApples(apples []int, days []int) int {
    /*
    第一步
    由于要挑出快过期的苹果先吃掉，为此，需要知道每批苹果的保质期内的lastday
        以lastday为key，这批苹果数量为value，构建哈希表:
        map[key:保质期最后一天lastday, value:下一天即将腐烂的苹果总数cnt]
    第二步
    先吃快过期的，即，我们要挑lastday小的先吃，同时，由于每天都有新苹果长出，所以每日会有新的lastday/cnt的加入
    为此借助堆，维护这个lastday的插入，构建小顶堆，队首的lastday一定最小
    */
    //map[key:保质期最后一天lastday, value:下一天即将腐烂的苹果总数cnt]
    var fre=make(map[int]int)
    var m=len(apples)
    var i int
    var minHeap=&IntHeap{}
    heap.Init(minHeap)//要挑lastday小的先吃。大顶堆，
    var res int
    for {
        if i<m{
            //m天内都有新苹果加入
            lastDay:=i+days[i]-1//保质期内最后一天的日期
            cnt:=apples[i]
            if cnt>0{
                heap.Push(minHeap,lastDay)
                fre[lastDay]+=cnt
            }
        }
        //吃苹果，挑快过期的先吃
        for minHeap.Len()>0{
            top:=minHeap.Top()
            if top<i||fre[top]==0{
                //已经过期，已经清空的
                heap.Pop(minHeap)
                continue
            }
            fre[top]-=1//吃掉最快过期的一个苹果
            res++
            break
        }
        //len天之后，就不会有新的苹果加入了，此后，坐吃山空
        if i>=m{
            if minHeap.Len()<1{
                break
            }
        }
        i++
    }
    return res
}
type IntHeap[]int
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int)bool {
    return h[i]<h[j]//min
}
func (h *IntHeap)Push(x interface{}) {
    *h=append(*h,x.(int))
}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
func (h IntHeap)Top()int {
    if len(h)<1 {
        return 0
    }
    return h[0]
}


