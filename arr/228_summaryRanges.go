func summaryRanges(nums []int) []string {
    var res[]string
    if len(nums)<1 {
        return res
    }
    var m=len(nums)
    var start=nums[0]
    var end=start
    for i:=0;i<m-1;i++{
        if nums[i]+1==nums[i+1] {
            end=nums[i+1]
        }else {
            startStr:=strconv.Itoa(start)
            endStr:=startStr
            finalStr:=startStr
            if start!=end {
                endStr=strconv.Itoa(end)
                finalStr=startStr+"->"+endStr
            }
            res=append(res,finalStr)
            //new
            start=nums[i+1]
            end=start
        }
    }
    startStr:=strconv.Itoa(start)
    endStr:=startStr
    finalStr:=startStr
    if start!=end {
        endStr=strconv.Itoa(end)
        finalStr=startStr+"->"+endStr
    }
    res=append(res,finalStr)
    return res
}
