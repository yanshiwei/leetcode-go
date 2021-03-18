func findBestValue(arr []int, target int) int {
    var sum int
    var n=len(arr)
    sort.Ints(arr)
    for i:=range arr{
        //计算平均值，如果从第i个开始所有元素为tmp,此时这些和为target-sum=(n-i)*tmp
        tmp:=float64(target-sum)/float64(n-i)
        if float64(arr[i])>tmp{
            //因为排过序了，如果arr[i]>tmp,后面的元素也必定>tmp
            return int(tmp+0.4)//返回整数，所以五舍六入 
        }
        sum+=arr[i]
    }
    return arr[n-1]//都没有找到，则直接使用最后一个
}
