func canThreePartsEqualSum(arr []int) bool {
    var sum int
    for i:=range arr {
        sum+=arr[i]
    }
    if sum%3!=0{
        return false
    }
    var avg=sum/3
    var segmentSum,i int
    for i=range arr {
        segmentSum+=arr[i]
        if segmentSum==avg {
            break
        }
    }
    if segmentSum!=avg{
        return false
    }
    var j=i+1
    for j+1<len(arr){
        // j+1需要满足最后一个数组非空
        segmentSum+=arr[j]
        if segmentSum==avg*2{
            return true//剩下的1/3不用再找了
        }
        j++
    }
    return false
}
