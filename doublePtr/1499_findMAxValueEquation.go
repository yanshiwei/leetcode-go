func findMaxValueOfEquation(points [][]int, k int) int {
    var res=INTMIN
    /*
    暴力 超时
    for i:=0;i<len(points)-1;i++{
        start:=points[i]
        for j:=i+1;j<len(points);j++{
            end:=points[j]
            if abs(start[0]-end[0])<=k{
                res=max(res,start[1]+end[1]+abs(start[0]-end[0]))
            }
        }
    }
    */
    /*
    滑动窗口，参考：https://leetcode-cn.com/problems/max-value-of-equation/solution/1499-java-hua-dong-chuang-kou-bai-fen-ba-ey7s/
    根据条件知道 1<=i<j<=points.length 可得 xi<xj，所以由 |xi-xj|<=k可得 xj-xi<=k,接着由 yi+yj+|xi-xj|可得 yi+yj+xj-xi，
    即对于(xi,yi)与（xj,yj）有yi+yj+xj-xi
    则对于（xi1,yi1）与（xj,yj）有yi1+yj+xj-xi1
    对于（xi0,yi0）与（xj,yj）有yi0+yj+xj-xi0
    固定j求其中较大值若yi1+yj-xj-xi1-(yi0+yj+xj-xi0)>0可得 yi1-xi1>yi0-xi0。解释：
    假设left和right的位置是left1和left2，也就是后面另有个right，判断right相同时left1、left2分别代入不等式，如果left2大就说明后面可能会有更大的，所以我们要对left进行移动，因为是可能会有更大的，所以就有yi2 + yj + xj - xi2 - (yi1 + yj + xj - xi1) > 0的情况下left移动，yj + xj前后抵消得出yi2 - xi2 > yi1 - xi1的情况下left移动
    */
    var left,right int
    for right<len(points){
        if left>=right{
            right++
            continue
        }
        //寻找xj-xi<=k
        if points[right][0]-points[left][0]>k{
            //left剔除，缩小窗口
            left++
            continue
        }
        //find max
        res=max(res,points[left][1]+points[right][1]+points[right][0]-points[left][0])
        // yi2-xi2>yi1-xi1
        if points[right][1]-points[right][0]>points[left][1]-points[left][0]{
            //说明right后面可能会有更大
            left++
            continue
        }
        right++
    }
    return res
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
