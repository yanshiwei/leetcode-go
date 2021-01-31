func majorityElement(nums []int) []int {
    var m=len(nums)
    var k=m/3
    var res[]int
    if len(nums)<1 {
        return res
    }
    var fre=make(map[int]int)
    for i:=range nums {
        if _,ok:=fre[nums[i]];ok==false {
            fre[nums[i]]=1
        }else {
            fre[nums[i]]+=1
        }
    }
    for key,v:=range fre {
        if v>k{
            res=append(res,key)
        }
    }
    return res
}
