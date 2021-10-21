//https://leetcode-cn.com/problems/gray-code/solution/gray-code-jing-xiang-fan-she-fa-by-jyd/
func grayCode(n int) []int {
    var res[]int
    res=append(res,0)
    if n==0{
        return res
    }
    var head=1
    for i:=0;i<n;i++{
        // n位gray码需要镜像构造n次
        for j:=len(res)-1;j>=0;j--{
            // G‘(n) 由G（n）前面加0得到，结果值相同无需构造；每次只需通过G（n）的逆序R(n),在前面+1构造R('n)
            res=append(res,head+res[j])
        }
        head<<=1//每次head后面多加个0: 1, 10 , 100 , 1000 ,....
    }
    return res
}
