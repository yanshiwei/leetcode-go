func moveZeroes(nums []int)  {
    if len(nums)<1 {
        return
    }
    var m=len(nums)
    var zoroLen int
    for i:=m-1;i>=0;i-- {
        if nums[i]==0 {
            if i==m-1 {
                continue
            }
            zoroLen++
            for j:=i;j<m-1;j++ {
                nums[j]=nums[j+1]
            }
            
        }
    }
    for i:=m-1;i>=0;i-- {
        if zoroLen>0 {
            nums[i]=0
            zoroLen--
        }
    }
    return
}
