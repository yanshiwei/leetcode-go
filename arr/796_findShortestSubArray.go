func findShortestSubArray(nums []int) int {
    var m=len(nums)
    var res int
    if m<1{
        return res
    }
    res=m+1
    var maxDegree int=0
    var fre=make(map[int]int)
    for i:=range nums{
        fre[nums[i]]++
        if fre[nums[i]]>maxDegree{
            maxDegree=fre[nums[i]]
        }
    }
    var first,last int
    for k,v:=range fre{
        if v==maxDegree{
            //找到元素出现次数最多的元素的下标
            for i:=0;i<m;i++{
                if nums[i]==k{
                    //第一个位置
                    first=i
                    break
                }
            }
            //找到最后一个元素
            for i:=m-1;i>=0;i--{
                if nums[i]==k{
                    last=i
                    break
                }
            }
            if last-first+1<res{
                res=last-first+1
            }
        }
    }
    return res
}
