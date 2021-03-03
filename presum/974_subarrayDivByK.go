func subarraysDivByK(A []int, K int) int {
    /*前缀和preSum[i]代表前i个元素和
    preSum[0]=0
    preSum[1]=preSum[0]+A[0]
    preSum[i]=preSum[i-1]+A[i-1]
    连续子数组的和就可以表示为前缀和的差,sum(A[i : j + 1]) = s[j + 1] - s[i]
    如果两个数的差能被K整除，就说明这两个数 mod K得到的结果相同: 证明(x-y)/k=a
    x=k*a+y,x%k=(k*a+y)%k=y%k
    只要找有多少对 mod k 相同的数就可以组合一下得到结果
    fre key 是preSum %k的值， value是个数
    对于A= [4,5,0,-2,-3,1] K = 5 ，preSum= [0, 4, 9, 9, 7, 4, 5]
    则fre[0]=2,fre[2]=1,fre[4]=4故在保证余数相同的情况下，取出两个数都可以得到一组答案。对于这个例子答案就是 C22 + C12 + C42 = 1 + 0 + 6 = 7
    C(n,m)=P(n,m)/P(m,m) =n!/m!（n-m）!,其中m=2
    则C(n,2)=n*(n-1)/2
    */
    var m=len(A)
    var preSum=make([]int,m+1)
    for i:=0;i<m;i++ {
        preSum[i+1]=preSum[i]+A[i]
    }
    var fre=make(map[int]int)
    for i:=range preSum{
        if preSum[i]<0&&preSum[i]%K!=0 {
            //先把负数转为正再计算余数 如a是负数 k-abs(a)%k
            preSum[i]=K-abs(preSum[i])%K
        }
        fre[preSum[i]%K]++
    }
    //fmt.Println(fre)
    var sum int
    for _,v:=range fre{
        sum+=v*(v-1)/2
    }
    return sum
}
func abs(a int)int {
    if a<0{
        return -1*a
    }
    return a
}
