func getMaximumGenerated(n int) int {
    if n==0{
        return 0
    }
    var arr=make([]int,n+1)
    var maxV =1
    for i:=range arr{
        if i<2 {
            arr[i]=i
        }else{
            if i%2==0{
                //偶数
                arr[i]=arr[i/2]
            }else{
                //奇数
                arr[i]=arr[i/2]+arr[i/2+1]
            }
            if arr[i]>maxV{
                maxV=arr[i]
            }
        }
    }
    return maxV
}
