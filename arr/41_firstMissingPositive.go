func firstMissingPositive(nums []int) int {
    /*
    总体思路：
    对于一个长度为 N 的数组，其中没有出现的最小正整数只能在 [1, N+1]中，因为：
    因为如果 [1, N]都出现了，那么答案是 N+1；否则是 [1, N]中没有出现的最小正整数。具体而言，
    对数组进行遍历，对于遍历到的数 x，如果它在 [1, N]的范围内，那么就将数组中的第 x-1个位置（注意：数组下标从 0开始）打上「标记」。在遍历结束之后，如果所有的位置都被打上了标记，那么答案是 N+1，否则答案是最小的没有打上标记的位置加1；
    问题的关键就是如何在常数空间下进行标记，由于我们只在意 [1, N]中的数，因此我们可以先对数组进行遍历，把不在 [1, N]范围内的数修改成任意一个大于 N 的数（例如 N+1）。这样一来，数组中的所有数就都是正数了，因此我们就可以将「标记」表示为「负号」
    算法流程：
    1.我们将数组中所有小于等于 0的数修改为 N+1；
    2.遍历数组中的每一个数 x，它可能已经被打了标记，因此原本对应的数为 |x|。如果∣x∣∈[1,N]，那么我们给数组中的第 |x| - 1个位置的数添加一个负号。注意如果它已经有负号，不需要重复添加
    3.遍历完成之后，如果数组中的每一个数都是负数，那么答案是 N+1，否则答案是第一个正数的位置加 1。
    */
    var n=len(nums)
    for i:=range nums{
        //step1
        if nums[i]<=0{
            nums[i]=n+1
        }
    }

    //所有数都是正的
    for i:=range nums{
        num:=abs(nums[i])//之所以绝对值，数可能已经打了标签
        //step2
        if num<=n{
            nums[num-1]=-1*abs(nums[num-1])
        }
    }
    //step3
    for i:=range nums{
        if nums[i]>0{
            return i+1
        }
    }
    return n+1
}

func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}