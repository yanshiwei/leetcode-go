func threeConsecutiveOdds(arr []int) bool {
    var m=len(arr)
    if m<3{
        return false
    }
    var  oddCnt int
    for i:=0;i<m;i++{
        if arr[i]%2==1{
            oddCnt++
            if oddCnt==3{
                return true
            }
        }else{
            oddCnt=0
        }
    }
    return false
}
