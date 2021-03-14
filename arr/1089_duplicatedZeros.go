func duplicateZeros(arr []int)  {
    if len(arr)<1 {
        return
    }
    var visi=make(map[int]struct{})
    for i:=range arr {
        //fmt.Println(arr,i,arr[i],visi)
        if _,ok:=visi[i];ok==true{
            continue
        }
        if arr[i]==0{
            if i+1>=len(arr){
                break
            }
            tmp:=append([]int(nil),arr[i+1:]...)
            if i+2>=len(arr){
                arr[i+1]=0
                visi[i+1]=struct{}{}
                break
            }
            curIdx:=0
            for j:=i+2;j<len(arr);j++{
                arr[j]=tmp[curIdx]
                curIdx++
            }
            arr[i+1]=0
            visi[i+1]=struct{}{}
        }
    }
    return
}
