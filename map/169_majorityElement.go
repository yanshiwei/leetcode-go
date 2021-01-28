func majorityElement(nums []int) int {
    var m=len(nums)
    var n=m/2
    var fre=make(map[int]int)
    for i:=range nums {
        if _,ok:=fre[nums[i]];ok==false {
            fre[nums[i]]=1
        }else {
            fre[nums[i]]+=1
        }
    }
    //fmt.Println(fre)
    for k,v:=range fre {
        if v>n {
            return k
        }
    }
    return 0
}
