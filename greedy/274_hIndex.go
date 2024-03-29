func hIndex(citations []int) int {
    /* 
* 思路：
*
* 1、首先看到h个元素大于等于某个值，N-h个元素小于等于某个值，这显然是一个有序序列的特征，所以自然而然的想到先将数组排序；
* 
* 2、将数组排序之后，对于给定的某个i，我们知道有citations.length - i篇论文的引用数 ≥ citations[i]，i篇
*    论文的引用数 ≤ citations[i]；
* 
* 3、不妨设h = citations.length - i，即至多有h篇论文分别引用了至少citation[i]次，其余citations.length - h篇
*    论文的引用数不多于citation[i]次。
*
*    既然如此，只要citation[i] ≥ h，就满足题意。
*/
    sort.Ints(citations)
    var h int
    for i:=0;i<len(citations);i++{
        h:=len(citations)-i
        if h<=citations[i]{
            return h
        }
    }
    return h
}
