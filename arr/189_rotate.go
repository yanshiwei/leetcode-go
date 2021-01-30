func rotate(nums []int, k int)  {
    /*
    当我们将数组的元素向右移动 k次后，尾部k%m 个元素会移动至数组头部，其余元素向后移动 k%m个位置
    类似效果：先将所有元素反转，这样尾部的k%m 个元素会移动至数组头部，然后再对[0,k%m-1]区间元素和[k%m-m-1]区间元素进行翻转，保证了顺序
    */
    var m=len(nums)
    if m<1 {
        return
    }
    //全部翻转
    reverse(nums,0,m-1)
    //[0,k%m-1]
    reverse(nums,0,k%m-1)
    //[k%m,m-1]
    reverse(nums,k%m,m-1)
}

func reverse(nums[]int,start,end int) {
    for start<end {
        nums[start],nums[end]=nums[end],nums[start]
        start++
        end--
    }
}


