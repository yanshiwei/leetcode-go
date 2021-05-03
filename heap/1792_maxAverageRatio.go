func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    /*
    贪心+堆
    1.先分别计算所有班级如果加一个人的通过变化率，放入堆。假设某班有 total 名学生，有 pass 名通过了考试，则该班级的通过率为 x = pass / total，该班级对年级通过率的影响为 x / n；若把一个聪明学生分配给该班级，则该加入学生对该班级通过变化率 deltaX = (pass+1) / (total + 1) - pass / total。
    2.按照贪心思想，每次分配时应该分配给产生贡献最大的班级（这里通过最大堆，每次取头部表示贡献最大），从而使最终的贡献最大；
    3.更新第二步的班级的新的通过变化率，随后将其重新加入（故一个班有可能被分配多次聪明学生）
    */
    var maxHeap=&Heap{}
    heap.Init(maxHeap)
    //计数每个聪明学生对每个班的提高因子
    for i:=range classes{
        oneClass:=classes[i]
        //原始通过率
        oldPassRatio:=float64(oneClass[0])/float64(oneClass[1])
        //加一个聪明学生后的通过率
        newPassRatio:=float64(oneClass[0]+1)/float64(oneClass[1]+1)
        heap.Push(maxHeap,Info{
            Idx:i,
            Score:newPassRatio-oldPassRatio,
        })
    }
    //开始真正分配聪明学生，每次分配时应该分配给产生贡献最大的班级，从而使最终的贡献最大，是一种贪心策略
    for extraStudents>0{
        biggestInfo:=maxHeap.Top()
        heap.Pop(maxHeap)
        //分配给biggestInfo.Idx班
        classes[biggestInfo.Idx][0]++
        classes[biggestInfo.Idx][1]++
        //原始通过率
        oldPassRatio:=float64(classes[biggestInfo.Idx][0])/float64(classes[biggestInfo.Idx][1])
        //加一个聪明学生后的通过率
        newPassRatio:=float64(classes[biggestInfo.Idx][0]+1)/float64(classes[biggestInfo.Idx][1]+1)
        biggestInfo.Score=newPassRatio-oldPassRatio
        heap.Push(maxHeap,biggestInfo)
        extraStudents--
    }
    var res float64
    for i:=range classes{
        oneClass:=classes[i]
        res+=float64(oneClass[0])/float64(oneClass[1])
    }
    return res/float64(len(classes))
}

type Info struct {
    Idx int
    Score float64
}

type Heap []Info
func (h Heap)Len()int{return len(h)}
func (h Heap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h Heap)Less(i,j int)bool{return h[i].Score>h[j].Score}//max
func (h Heap)Top()Info{return h[0]}
func (h *Heap)Push(x interface{}){*h=append(*h,x.(Info))}
func (h *Heap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
