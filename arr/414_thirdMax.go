func thirdMax(nums []int) int {
    /*
    每次更新第一大值和第二大值时，都传递给第三大值
    注意重复的跳过，最后判断第三大有没有更新过
    */
    var res int
    if len(nums)==1 {
        return nums[0]
    }
    if len(nums)==2 {
        if nums[0]>nums[1] {
            return nums[0]
        }
        return nums[1]
    }
    var firstMax=INTMIN
    var secondMax=firstMax
    res=firstMax
    for i:=0;i<len(nums);i++ {
        if firstMax<nums[i] {
            res=secondMax
            secondMax=firstMax
            firstMax=nums[i]
        }else if firstMax==nums[i] {
            continue
        }else if secondMax<nums[i] {
            res=secondMax
            secondMax=nums[i]
        }else if secondMax==nums[i] {
            continue
        }else if res<nums[i] {
            res=nums[i]
        }
    }
    if res==INTMIN {
        //没更新过,返回最大值
        return firstMax
    }
    return res
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
