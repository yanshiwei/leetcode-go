//埃氏筛法的原理是：如果一个数是素数，那么这个数的倍数一定不是素数
func countPrimes(n int) int {
    // var cnt int
    // for i:=2;i<n;i++{
    //     if isPrime(i){
    //         cnt++
    //     }
    // }
    // return cnt
    if n<=2{
        return 0
    }
    var flag=make([]bool,5000000)
    var cnt int
    for i:=2;i<n;i++{
        //如果第一次遍历到，则它是素数
        if flag[i]==false{
            cnt++
            for j:=i;j<n;j+=i{
                flag[j]=true
            }
        }
    }
    return cnt
}
//某个数字 n，i 从 2 开始枚举判断是否满足n%i==0 ，如果找到了 n 的因子，就返回 false。注意 i遍历到最大sqrt(n)即可
func isPrime(num int)bool{
    var maxNum=int(math.Sqrt(float64(num)))
    for i:=2;i<=maxNum;i++{
        if num%i==0{
            return false
        }
    }
    return true
}
