func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    var m1=len(arr1)
    var m2=len(arr2)
    var res int
    for i:=0;i<m1;i++{
        isFound:=true
        for j:=0;j<m2;j++{
            if abs(arr1[i]-arr2[j])<=d{
                isFound=false
                break
            }
        }
        if isFound{
            res++
        }
    }
    return res
}

func abs(x int)int {
    if x<0{
        return -1*x
    }
    return x
}
