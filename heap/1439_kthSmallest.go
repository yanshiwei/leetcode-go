func kthSmallest(mat [][]int, k int) int {
    /*
    题目从每一行必去一个，形成数组求和，要去和最小的第K个。
    每次从堆中弹出最小的数组和curr_sum和对应的指针pointers，然后轮流将指针pointers的每个索引向后移动一位，生成新的new_sum，加入堆中。
    1.最小堆存储的是[curr_sum, pointers]二元组，pointers是指针数组，curr_sum是该pointers指向的元素的和。初始化pointers全为0，求出相应的curr_sum，并将其入堆。即初始化堆为第一列的元素和，并且保存位置[0, 0,..., 0]
    2.重复下列步骤k次
        2.1 从堆中pop出curr_sum和pointers
        2.2 遍历pointers的每个索引，将该索引加一，求出新的和，如果没有出现过，push入堆
    1 3 11
    2 4 6
    类似链表：
    1->3->11
    2->4->6
    第一列最小，后续每一行向后移动一位，每一行均入堆，则肯定右产生一个最值，如此操作K次就得到第K个最值
    举例：
    假设有M=2的一下矩阵

​   [a1,a2,a3]

​   [b1,b2,b3]

​   根据题意，每行以非递减排序，则有：

​   a1+b1 <= a1+b2 <= a1+b3

​   a2+b1 <= a2+b2 <= a2+b3

​   a3+b1 <= a3+b2 <= a3+b3
    */
    var m=len(mat)
    var n=len(mat[0])
    if m<1||n<1{
        return -1
    }
    var pointers string//[0,0,0]每一行的第一个
    for i:=0;i<m-1;i++{
        pointers+="0,"
    }
    pointers+="0"
    var minHeap=&IntHeap{}
    heap.Init(minHeap)
    var curSum int
    for i:=range mat{
        curSum+=mat[i][0]//将每一行第一列加入
    }
    heap.Push(minHeap,Info{CurSum:curSum,CurPointer:pointers})
    var visited=make(map[string]bool)
    visited[pointers]=true//该数组有没有出现过
    for i:=0;i<k;i++{
        //从堆中pop出curr_sum(最小数组和)和对应的下标数组pointer
        if minHeap.Len()<1{
            break
        }
        info:=heap.Pop(minHeap).(Info)
        curSum,pointers=info.CurSum,info.CurPointer
        pointersList:=strings.Split(pointers,",")
        //每个指针轮流后移一位，将new_sum(新的数组和)和new_pointers(新的指针数组)push入堆
        for j:=range pointersList{
            //j代表mat的行数
            curIdxStr:=pointersList[j]
            curIdx,_:=strconv.Atoi(curIdxStr)
            //curIdx代表mat的第j行第curIdx位置
            if curIdx<n-1{
                var newPointersList []string
                for k:=range pointersList{
                    newPointersList=append(newPointersList,pointersList[k])
                }
                newPointersList[j]=strconv.Itoa(curIdx+1)//后移一位
                newPointersStr:=strings.Join(newPointersList,",")
                if _,ok:=visited[newPointersStr];ok==false{
                    newSum:=curSum+mat[j][curIdx+1]-mat[j][curIdx]
                    heap.Push(minHeap,Info{CurSum:newSum,CurPointer:newPointersStr})
                    visited[newPointersStr]=true
                }
            }
        }
    }
    return curSum
}
type Info struct {
    CurPointer string
    CurSum int
}
type IntHeap []Info
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){
    if i<len(h)&&j<len(h){
        h[i],h[j]=h[j],h[i]
    }
}
func (h IntHeap)Less(i,j int)bool {
    return h[i].CurSum<h[j].CurSum//min
}
func (h *IntHeap)Push(x interface{}){
    *h=append(*h,x.(Info))
}
func (h IntHeap)Top()Info {
    return h[0]
}
func (h *IntHeap)Pop()interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
