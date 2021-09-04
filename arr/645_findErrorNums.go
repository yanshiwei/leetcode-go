func findErrorNums(nums []int) []int {
    var res []int
    if len(nums)<2{
        return res
    }
    var n=len(nums)
    var fre=make([]int,n+1)
    var dup int
    for i:=0;i<len(nums);i++{
        if fre[nums[i]]==0{
            fre[nums[i]]=nums[i]
        }else{
            dup=nums[i]
        }
    }
    //fmt.Println(fre)
    for i:=1;i<len(fre);i++{
        if fre[i]==0{
            res=append(res,dup)
            res=append(res,i)
            return res
        }
    }
    return res
}
