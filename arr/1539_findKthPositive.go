func findKthPositive(arr []int, k int) int {
    var cnt=arr[0]-1
    var missIdx=arr[0]
    if k<arr[0]{
        return k
    }
    var i=0
    for i<len(arr){
        if arr[i]!=missIdx{
            cnt++
            if cnt>=k{
                return missIdx
            }
            missIdx++
            continue
        }
        i++
        missIdx++
    }
    //fmt.Println(k,i,missIdx)
    k=k-cnt
    missIdx--
    for k>0{
        missIdx++
        k--
    }
    return missIdx
}

const INTMAx=int(^uint(0)>>1)
