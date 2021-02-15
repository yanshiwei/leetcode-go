func dominantIndex(nums []int) int {
    var first =INTMIN
    var firstV int=-1
    for i:=range nums{
        if nums[i]>first{
            firstV=i
            first=nums[i]
        }
    }
    //fmt.Println(firstV,first)
    for i:=range nums{
        if i!=firstV&&nums[i]*2>first{
            return -1
        }
    }
    return firstV
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
