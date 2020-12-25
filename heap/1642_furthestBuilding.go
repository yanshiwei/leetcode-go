func furthestBuilding(heights []int, bricks int, ladders int) int {
    //基本思想 在高度差较大时，尽量用梯子，因为梯子无限长，而砖头有限，尽量节省
    //先假设梯子无限个，先都使用梯子，直到梯子数目不够，则寻找高度差最小的将梯子换成转头
    //自己实现的小顶堆总是超时，只能借助于container/heap
    //var ç=NewHeap(len(heights))
    minHeap:=&IntHeap{}
    heap.Init(minHeap)//container/heap
    for i:=1;i<len(heights);i++{
        delta:=heights[i]-heights[i-1]
        if delta<=0{
            //h[i]<=h[i-1]无需任何操作
            continue
        }
        //minHeap.Push(delta)
        heap.Push(minHeap,delta)
        if minHeap.Len()<=ladders{
            //先默认都是用梯子
            continue
        }else {
            //梯子不够 找之前最小高度差，才使用砖头，以节省砖头
            //minV:=minHeap.Pop()
            minV:=heap.Pop(minHeap).(int)
            bricks-=minV
            if bricks<0{
                //=0 需继续执行
                return i-1
            }
        }
    }
    return len(heights)-1//整个height都能到达
}

func adjustHeap(minHeap[]int,start,length int){
    var i=start
    var child  int
    for {
        child=2*i+1
        if child>length-1{
            break
        }
        if child+1<=length-1&&minHeap[child]>minHeap[child+1]{
            child++
        }
        if minHeap[i]>minHeap[child]{
            minHeap[i],minHeap[child]=minHeap[child],minHeap[i]
            i=child
        }else {
            break
        }
    }
}

func buildHeap(minHeap[]int){
    for i:=len(minHeap)/2-1;i>=0;i--{
        adjustHeap(minHeap,i,len(minHeap))
    }
}
type Heap struct {
	Size  int
	Elems []int
}

func NewHeap(MaxSize int) *Heap {
	h := new(Heap)
	h.Elems = make([]int, MaxSize, MaxSize)
	return h
}

func (h *Heap) Push(x int) {
	h.Size++

	// i是要插入节点的下标
	i := h.Size-1
	for {
		if i <= 0 {
			break
		}

		// parent为父亲节点的下标
		parent := (i - 1) / 2
		// 如果父亲节点小于等于插入的值，则说明大小没有颠倒，可以退出
		if h.Elems[parent] <= x {
			break
		}

		// 互换当前父亲节点与要插入的值
		h.Elems[i] = h.Elems[parent]
		i = parent
	}

	h.Elems[i] = x
    fmt.Println("push",h.Elems)
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return 0
	}

	// 取出根节点
	ret := h.Elems[0]

	// 将最后一个节点的值提到根节点上
	h.Size--
	x := h.Elems[h.Size]

	i := 0
	for {
		// a，b为左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 没有左子树
		if a >= h.Size {
			break
		}

		// 有右子树，找两个子节点中较小的值
		if b < h.Size && h.Elems[b] < h.Elems[a] {
			a = b
		}

		// 父亲小直接退出
		if h.Elems[a] >= x {
			break
		}

		// 交换
		h.Elems[i] = h.Elems[a]
		i = a
	}

	h.Elems[i] = x
    fmt.Println("pop",h.Elems)
	return ret
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

