func maxValueAfterReverse(nums []int) int {
    /*
    贪心
    题意：nums = [2,5,1,3,4]
    数组和为
    abs(a[0]-a[1])+abs(a[1]-a[2])+abs(a[2]-a[3])+abs(a[3]-a[4])
    =3+4+2+1=10
    我们假设反转[l,r]:
原数组 ... a[l - 1] ,a[l] , ... , a[r] , a[r + 1] ...
反转后 ... a[l - 1] , a[r] , ... , a[l] , a[r + 1]...
通过计算发现发生改变的数值仅仅是两个端点,如[2,3,1,5,4] 改变翻转子数组 [3,1,5] ，数组变成 [2,5,1,3,4]，子数组[3,1,5]内部与[5,1,3]数组值不变仍为6
    减少的值 = abs(a[l] - a[l - 1]) + abs(a[r] - a[r + 1]);
    增加的值 = abs(a[l - 1] - a[r]) + abs(a[l] - a[r + 1]);
    变化的值 = abs(a[l - 1] - a[r]) + abs(a[l] - a[r + 1]) - (abs(a[l] - a[l - 1]) + abs(a[r] - a[r + 1]));
    枚举所有情况，增加的值=abs(a[l−1]−a[r])+abs(a[l]−a[r+1]) =
max{
    a[l - 1] + a[l] - (a[r] + a[r+1]) ,a[l−1]+a[l]−(a[r]+a[r+1]),
    a[l - 1] - a[l] - (a[r] - a[r + 1]),a[l−1]−a[l]−(a[r]−a[r+1]),
    -a[l - 1] + a[l] - (-a[r] + a[r + 1]),−a[l−1]+a[l]−(−a[r]+a[r+1]),
    -a[l - 1] - a[l] - (-a[r] - a[r + 1])−a[l−1]−a[l]−(−a[r]−a[r+1])
}
最后把变化的值计算出来
max{
    a[l - 1] + a[l] - abs(a[l] - a[l - 1]) - (a[r] + a[r + 1] + abs(a[r] - a[r + 1])),a[l−1]+a[l]−abs(a[l]−a[l−1])−(a[r]+a[r+1]+abs(a[r]−a[r+1])),
    a[l - 1] - a[l] - abs(a[l] - a[l - 1]) - (a[r] - a[r + 1] + abs(a[r] - a[r + 1])),a[l−1]−a[l]−abs(a[l]−a[l−1])−(a[r]−a[r+1]+abs(a[r]−a[r+1])),
    -a[l - 1] + a[l] - abs(a[l] - a[l - 1]) - (-a[r] + a[r + 1] + abs(a[r] - a[r + 1])),−a[l−1]+a[l]−abs(a[l]−a[l−1])−(−a[r]+a[r+1]+abs(a[r]−a[r+1])),
    -a[l - 1] - a[l] - abs(a[l] - a[l - 1]) - (-a[r] - a[r + 1] + abs(a[r] - a[r + 1]))−a[l−1]−a[l]−abs(a[l]−a[l−1])−(−a[r]−a[r+1]+abs(a[r]−a[r+1]))
    最后注意边界的处理
}
    */
    var m=len(nums)
    var resOri int
    for i:=1;i<m;i++{
        resOri+=abs(nums[i]-nums[i-1])
    }
    var res =INTMIN
    //边界处理 l==0时,原来的值abs(nums[r]-nums[i+1]);新值abs(nums[0]-nums[r+1])
    for r:=1;r<m-1;r++{
        delta:=abs(nums[0]-nums[r+1])-abs(nums[r]-nums[r+1])
        res=max(res,resOri+delta)
    }
    //r=n-1边界，原来的值abs(nums[l]-nums[l-1]), 新值abs(nums[l-1]-nums[m-1])
    for l:=1;l<m-1;l++{
        delta:=abs(nums[l-1]-nums[m-1])-abs(nums[l]-nums[l-1])
        res=max(res,resOri+delta)
    }
    //普通处理，
    //原来的值abs(nums[l]-nums[l-1])+abs(nums[r]-nums[i+1])
    //新的值 abs(nums[l-1]-nums[m-1])+abs(nums[l]-nums[r+1])
    var f1,f2,f3,f4=INTMIN,INTMIN,INTMIN,INTMIN
    for l:=1;l<m;l++{
        absV:=abs(nums[l-1]-nums[l])
        f1=max(f1,nums[l-1]+nums[l]-absV)
        f2=max(f2,nums[l-1]-nums[l]-absV)
        f3=max(f3,-nums[l-1]+nums[l]-absV)
        f4=max(f4,-nums[l-1]-nums[l]-absV)
    }
    var g1,g2,g3,g4=INTMIN,INTMIN,INTMIN,INTMIN
    for r:=0;r<m-1;r++{
        absV:=abs(nums[r]-nums[r+1])
        g1=max(g1,-nums[r]-nums[r+1]-absV)
        g2=max(g2,-nums[r]+nums[r+1]-absV)
        g3=max(g3,nums[r]-nums[r+1]-absV)
        g4=max(g4,nums[r]+nums[r+1]-absV)
    }
    //fi gi注意一一对应
    fmt.Println(f1,f2,f3,f4,resOri,res)
    fmt.Println(g1,g2,g3,f4)
    delta:=getMAxFromArr([]int{f1+g1,f2+g2,f3+g3,f4+g4})
    return max(res,resOri+delta)
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX

func getMAxFromArr(a []int)int{
    var res=INTMIN
    for i:=range a{
        if res<a[i]{
            res=a[i]
        }
    }
    return res
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
func abs(z int)int {
    if z<0{
        z=-1*z
    }
    return z
}
