func canReorderDoubled(arr []int) bool {
    /*
    map统计每个数出现次数
    从小到大排序，对于每个数（记为x），看看是否有相应数量的另一个数（记为y）与之匹配。因为是从小到大扫描，所以x < 0时候， y = x / 2； x > 0时候， y = x * 2。注意：如果负数个数不为2的倍数，肯定不对
    从最大到最小值遍历
    */
    var m=len(arr)
    var cntNegetive int
    var fre=make(map[int]int,m)
    for i:=0;i<m;i++ {
        if arr[i]<0{
            cntNegetive++
        }
        fre[arr[i]]++
    }
    if cntNegetive%2!=0{
        return false
    }
    sort.Ints(arr)
    for i:=range arr{
        //根据当前数字是否是负数，决定下一个值是乘以2还是除以2获得的值
        if fre[arr[i]]>0{
            fre[arr[i]]--
            if arr[i]<0&&arr[i]%2==-1{
                //负数值会涉及到/2. 当负数为奇数，直接/2会四舍五入，故case对待
                return false
            }
            var nextV int
            if arr[i]>0{
                nextV=arr[i]*2
            }else{
                nextV=arr[i]/2
            }
            if fre[nextV]<=0{
                //没有配对
                return false
            }
            fre[nextV]--
        }
    }
    return true
}
