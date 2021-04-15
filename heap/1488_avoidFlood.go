func avoidFlood(rains []int) []int {
    /*
    贪心+小顶堆 参考：https://leetcode-cn.com/problems/avoid-flood-in-the-city/solution/1488bi-mian-hong-shui-fan-lan-python3-tan-xin-li-y/
    https://leetcode-cn.com/problems/avoid-flood-in-the-city/solution/tan-xin-si-xiang-you-xian-dui-lie-nlogn-by-vector_/
    对于第一次下雨不理，倘若一个地方下两次雨，则必须在两次下雨之间有一个晴天把水抽干。
    当有一天不下雨时，我们可以选择一个湖泊来抽干它，那么最优情况应该怎么选择呢？为了避免任意一个湖泊发生洪水，所以应该选择这天之前，所有装满水的湖泊中，在这天之后最近的将会下雨的湖泊
    */
    var res =make([]int,len(rains))
    for i:=range res{
        res[i]=1
    }
    h:=&hp{}//小顶堆持续刷新当前已经满的湖后续最早下雨的时间
    heap.Init(h)
    n:=len(rains)
    nextRainDay:=make([]int,n)// 标记下一次下雨的日期
    rainLastDay:=make(map[int]int)//湖对应上一个下雨日期  key is 湖 value is 日期
    for i:=range rains{
        if rains[i]>0{
            //下雨
            if j,ok:=rainLastDay[rains[i]];ok{
                nextRainDay[j]=i
            }
        }
        nextRainDay[i]=n+1//晴天默认下一层日期
        rainLastDay[rains[i]]=i
    }
    //fmt.Println(nextRainDay)
    isFull:=make(map[int]bool)//key is 湖
    for i:= range rains {
        if rains[i]>0{
            res[i]=-1
            if isFull[rains[i]]{
                //满了
                return []int{}
            }
            //按下一次 最早下雨的湖泊 排序，（下一次下雨日期， 当前湖泊编号）
            h.push([]int{nextRainDay[i],rains[i]})
            isFull[rains[i]]=true
        }else{
            if len(h.days)>0{
                //如不下雨，则选取当前已经满的湖且后续最早下雨的湖进行抽水
                // 弹出下一次 最早下雨的湖泊编号，优先抽干最先发生洪水的湖
                dry:=h.pop()
                res[i]=dry[1]//湖
                isFull[dry[1]]=false
            }
        }
    }
    return res
}

type hp struct {
    days [][] int//下一次下雨日期， 当前湖泊编号
}
//小顶堆
func (h *hp)Less(i,j int)bool {
    //按照日期从大到小
    return h.days[i][0]<h.days[j][0]
}
func (h *hp)Len()int {
    return len(h.days)
}

func (h *hp)Swap(i, j int){
    h.days[i],h.days[j]=h.days[j],h.days[i]
}

func (h *hp)Push(v interface{}){
    h.days=append(h.days,v.([]int))
}

func (h *hp)Pop()interface{}{
    old:=h.days
    v:=old[len(old)-1]
    h.days=old[0:len(old)-1]
    return v
}

func (h *hp) push(v []int) { heap.Push(h, v) }
func (h *hp) pop() []int   { return heap.Pop(h).([]int) }
