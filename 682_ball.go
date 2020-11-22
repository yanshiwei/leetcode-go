func calPoints(ops []string) int {
    var sum int
    var arr []int
    if len(ops)<1 {
        return sum
    }
    for _,one:=range ops {
        if one=="+" {
            add:=arr[len(arr)-1]+arr[len(arr)-2]//comfirmed
            arr=append(arr,add)
        }else if one =="D" {
            dou:=arr[len(arr)-1]*2//confirmed
            arr=append(arr,dou)
        }else if one =="C" {
            arr=arr[0:len(arr)-1]//confirmed
        }else {
            intV,_:=strconv.Atoi(one)//confirmed must be number
            arr=append(arr,intV)
        }
    }
    for _,v:=range arr {
        sum+=v
    }
    return sum
}
