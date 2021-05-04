type MKAverage struct {
    MinHeapForMaxPart *IntHeap//max k element as maxPart
    MaxHeapForMinPart *IntHeap//min k element as minPart
    MinHeapForMidPart *IntHeap//max (m-k) element as maxMidPart
    MaxHeapForMidPart *IntHeap//min (m-k) element as minMidPart
    M int
    K int
    Sum int//sum of last m element
    SumForMinPart int//sum of minPart
    SumForMaxPart int//sum of maxPart
    delForMinPart map[int]int//to be deleted in min+mid
    delForMaxPart map[int]int//to be deleted in max+mid
    CntForMinPart int//min k element remove cnt in minPart
    CntForMaxPart int//max k element remove cnt in maxPart
    Buf []int//last m element
}

/*
    堆
    参考：https://leetcode-cn.com/problems/finding-mk-average/solution/jian-dan-yi-dong-de-dui-jie-fa-dong-tai-si0yu/
*/

func Constructor(m int, k int) MKAverage {
    var obje=MKAverage{
        MaxHeapForMinPart:&IntHeap{Mode:false},
        MinHeapForMaxPart:&IntHeap{Mode:true},
        MaxHeapForMidPart:&IntHeap{Mode:false},
        MinHeapForMidPart:&IntHeap{Mode:true},
        M:m,
        K:k,
        delForMinPart:make(map[int]int),
        delForMaxPart:make(map[int]int),
    }
    heap.Init(obje.MaxHeapForMinPart)
    heap.Init(obje.MaxHeapForMidPart)
    heap.Init(obje.MinHeapForMaxPart)
    heap.Init(obje.MinHeapForMidPart)
    return obje
}


func (this *MKAverage) AddElement(num int)  {
    if len(this.Buf)==this.M{
        //buf元素满了，移除一个
        num:=this.Buf[0]
        this.Buf=this.Buf[1:]//pop first
        this.Sum-=num
        this.delForMinPart[num]++
        if num<=this.MaxHeapForMinPart.Top(){
            // 小于最小k个数的最大值，则要从最小k队列中删除这个数
            this.SumForMinPart-=num
            this.CntForMinPart++
        }
        this.delForMaxPart[num]++
        if num>=this.MinHeapForMaxPart.Top(){
            // 大于最大k个数的最大值，则要从最大k队列中删除这个数
            this.SumForMaxPart-=num
            this.CntForMaxPart++
        }
    }
    this.Buf=append(this.Buf,num)
    this.Sum+=num
    //+插入时，如果 num 比 max 最小值大，肯定要加入 max 区。如果 num 比 min 最大值小，肯定要加入 min 区。否则一律加入 mid 区
    if this.MaxHeapForMinPart.Len()==0||num<=this.MaxHeapForMinPart.Top(){
        //新增数字小于等于最小k个数的最大数，入最小k队列 minPart
        heap.Push(this.MaxHeapForMinPart,num)
        this.SumForMinPart+=num
    }else{
        //否则如最大的m-k队列,maxMidPart
        heap.Push(this.MinHeapForMidPart,num)
    }
    if this.MinHeapForMaxPart.Len()==0||num>=this.MinHeapForMaxPart.Top(){
        //新增数字大于等于最大k个数的最小数，入最大k队列 maxPart
        heap.Push(this.MinHeapForMaxPart,num)
        this.SumForMaxPart+=num
    }else{
        //否则如最小的m-k队列,minMidPart
        heap.Push(this.MaxHeapForMidPart,num)
    }
    this.balanceHeapA()
	this.balanceHeapB()
}


func (this *MKAverage) CalculateMKAverage() int {
    if len(this.Buf)<this.M{
        return -1
    }
    return (this.Sum-this.SumForMinPart-this.SumForMaxPart)/(this.M-this.K*2)
}

func (this *MKAverage) balanceHeapA(){
    for{
        //min k ele
        for this.MaxHeapForMinPart.Len()>0&&this.delForMinPart[this.MaxHeapForMinPart.Top()]>0{
            //删除最小k中已标记未删除的数
            this.delForMinPart[heap.Pop(this.MaxHeapForMinPart).(int)]--
            this.CntForMinPart--
        }
        //最大的m-k ele
        for this.MinHeapForMidPart.Len()>0&&this.delForMinPart[this.MinHeapForMidPart.Top()]>0{
            //删除最大m-k中已标记删除的数
            this.delForMinPart[heap.Pop(this.MinHeapForMidPart).(int)]--
        }
        // 平衡最小k和最大m-k/mix and mid
        if this.MaxHeapForMinPart.Len()-this.CntForMinPart<this.K&&this.MinHeapForMidPart.Len()>0{
            //+mid 元素过多时，如果 min 数量不到 k，则从 mid 的小根堆取堆顶移动到 min 即可
            num:=heap.Pop(this.MinHeapForMidPart).(int)
            this.SumForMinPart+=num
            heap.Push(this.MaxHeapForMinPart,num)
        }else if this.MaxHeapForMinPart.Len()-this.CntForMinPart>this.K{
            //+min 元素过多时，由于 min 是最大堆，堆顶元素移动到 mid 部分即可。
            num:=heap.Pop(this.MaxHeapForMinPart).(int)
            this.SumForMinPart-=num
            heap.Push(this.MinHeapForMidPart,num)
        }else{
            break
        }
    }
}

func (this *MKAverage) balanceHeapB(){
    for {
        // 删除最大k中已标记未删除的数
        for this.MinHeapForMaxPart.Len()>0&&this.delForMaxPart[this.MinHeapForMaxPart.Top()]>0{
            this.delForMaxPart[heap.Pop(this.MinHeapForMaxPart).(int)]--
            this.CntForMaxPart--
        }
        //最小的m-k
        for this.MaxHeapForMidPart.Len()>0&&this.delForMaxPart[this.MaxHeapForMidPart.Top()]>0{
            this.delForMaxPart[heap.Pop(this.MaxHeapForMidPart).(int)]--
        }
        // 平衡最大k和最小m-k
        if this.MinHeapForMaxPart.Len()-this.CntForMaxPart<this.K&&this.MaxHeapForMidPart.Len()>0{
            //+mid 元素过多时，如果 max 数量不到 k，则从 mid 的大根堆取堆顶移动到 max 即可
            num:=heap.Pop(this.MaxHeapForMidPart).(int)
            this.SumForMaxPart+=num
            heap.Push(this.MinHeapForMaxPart,num)
        }else if this.MinHeapForMaxPart.Len()-this.CntForMaxPart>this.K{
            //+max 元素过多时，由于 max 是最小堆，堆顶元素移动到 mid 部分即可
            num:=heap.Pop(this.MinHeapForMaxPart).(int)
            this.SumForMaxPart-=num
            heap.Push(this.MaxHeapForMidPart,num)
        }else{
            break
        }
    }
}

type IntHeap struct{
    Heap []int
    Mode bool
}
func (h IntHeap)Len()int{return len(h.Heap)}
func (h IntHeap)Swap(i,j int){h.Heap[i],h.Heap[j]=h.Heap[j],h.Heap[i]}
func (h IntHeap)Less(i,j int)bool{
    if h.Mode{
        //true是小根堆
        return h.Heap[i]<h.Heap[j]
    }else{
        //false是大根堆
        return h.Heap[i]>h.Heap[j]
    }
}
func (h IntHeap)Top()int {
	if h.Len()>0 {
		return h.Heap[0]
	}
	return 0
}
func (h *IntHeap)Push(x interface{}) {
	h.Heap=append(h.Heap,x.(int))
}
func (h *IntHeap)Pop()interface{}{
	old:=h
	n:=len(old.Heap)
	x:=old.Heap[n-1]
	h.Heap=old.Heap[0:n-1]
	return x
}
/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
