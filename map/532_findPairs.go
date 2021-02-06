func findPairs(nums []int, k int) int {
    if k<0{
        return 0
    }
    var fre=make(map[int]int)
    for i:=0;i<len(nums);i++ {
        if _,ok:=fre[nums[i]];ok==false {
            fre[nums[i]]=0
        }else {
            fre[nums[i]]+=1
        }
    }
    var res int
    for key,value:=range fre {
        if k==0 {
            //只需要判断有没有重复数字
            if value>=1 {
                res++
            }
        }else {
            //k>0
            if _,ok:=fre[k+key];ok==true {
                res++
            }
        }
    }
    return res
}
