func findContentChildren(g []int, s []int) int {
    //贪心算法，先排序，最小的饭量喂饱最小的饿肚子孩子
    sort.Ints(g)
    sort.Ints(s)
    var i,j int
    for i<len(g)&&j<len(s) {
        if s[j]>=g[i] {
            //喂饱一个小孩
            i++
        }
        j++
    }
    return i
}
