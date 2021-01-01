func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
/*
应选取多个中【工资/质量】比值最大的质量来作为比例的基准，证明过程：
实际上我们先证明两个数，再推广到多个数，
举例：
干活质量分别为（a，b），对应工资为（A，B），不妨假设工资/质量比值（A/a >= B/b），
（1）如果我们选取比值较大的A/a对应得质量a来作为工资等比例的基准，那么B实际发放工资=A * b/a,由于前提条件工资/质量比值大小关系（A/a >= B/b 等价于 A * b >= B * a），则B实际发放工资=A * b/a>=(B * a)/a=B 满足对B最低工资的要去
（2）如果选择比值小的作基准，则A实际发放工资=B*a/b,由于前提条件工资/质量比值大小关系（A/a >= B/b），则A实际发放工资=B*a/b=B/b*a<=(A/a)*a=A 不满足A的最低工资要去。
同理，推广到多个数的情况下，我们应该选择【工资/质量】比值最大的那个质量来约分。
入对于三个元素：
干活质量分别为（a，b，c），对应工资为（A，B，C），假设【A/a >= B/b >= C/c】，将发放总工资 = A * 1 + A * b/a + A * c/a = A * a/a + A * b/a + A * c/a = (a + b + c) * A/a，那么此时，该公式即为本题的核心解题公式了。
故有：
1.预处理，按R=【工资/质量】比值进行排序，从小到大；
2.在选定第i(i >= K)个为基准值时，需要在前i-1个价值比中选取最低成本的前K-1个作为最终的选择结果。
其实，我们在选择第i个为基准时，其实明确了一个价值比，因为已经排序，则i之前R都比i的要小，那么对于i 之前的其他选择的K-1人也是应用i这个价值比来计算实际工资，实际工资的计算为<价值比 x 质量>，按照核心公式以及i的值已经是确定的，那么问题就转换为求前K-1个质量最低的和，可以维护一个大顶堆，堆内元素个数维持为K个，每次遍历计算是否更新堆顶值，同时更新最低支出(基准值 x 当前质量和)

即维护一个大小K的大顶堆，动态调整堆获取TOPK小的值，比较的对象是质量
*/
var ratioList RatioSli
for i:=range wage {
    r:=Info{Ratio:float64(wage[i])/float64(quality[i]),Quality:quality[i]}
    ratioList=append(ratioList,r)
}
sort.Sort(ratioList)//从小到大
var maxHeap=&IntHeap{}
heap.Init(maxHeap)
var totalQuality int
for i:=0;i<K;i++ {
    heap.Push(maxHeap,ratioList[i].Quality)
    totalQuality+=ratioList[i].Quality
}
var minCost=ratioList[K-1].Ratio*float64(totalQuality)//核心公式
// 选择其中第i(i >= K-1)名之后的工人最为最低工资，则在前i名中选择K-1名最低工资组合即可
for i:=K;i<len(ratioList);i++ {
    newRatio:=ratioList[i].Ratio
    if ratioList[i].Quality>=maxHeap.Top() {
        continue
    }
    beforeQuality:=heap.Pop(maxHeap).(int)
    totalQuality-=beforeQuality//减去
    heap.Push(maxHeap,ratioList[i].Quality)
    totalQuality+=ratioList[i].Quality
    newCost:=newRatio*float64(totalQuality)
    minCost=min(minCost,newCost)
}
return minCost
}
func min(x,y float64)float64 {
    if x-y>1e-9 {
        return y
    }
    return x
}
type Info struct {
    Ratio float64
    Quality int
}
type RatioSli[]Info//max heap for quality
func (h RatioSli)Len()int {return len(h)}
func (h RatioSli)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h RatioSli)Less(i,j int)bool {
    return h[i].Ratio<h[j].Ratio
}

type IntHeap[]int//max heap for quality
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int)bool {
    return h[i]>h[j]//max
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
