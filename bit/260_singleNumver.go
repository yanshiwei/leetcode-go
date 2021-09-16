func singleNumber(nums []int) []int {
    var bitmask int
    for i:=range nums{
        bitmask^=nums[i]
    }
    //bitmask=a^b
    //在二进制中可能会有多个1，为了方便计算
    //我们只需保留其中的任何一个1，其他的都
    //让他变为0，这里保留的是最右边的1
    var div=1
    for (div&bitmask)==0{
        div<<=1
    }
    //对每一个数进行分组，将div这一位数值分别为0和1的数字分为两组（相同数值每一位都相同，会分到同一组，而唯一的两个数字这一位不同，会被分到两组）,那么每组就变成了只有一个数字出现一次，其他数字都出现两次，变成了136题，https://leetcode-cn.com/problems/single-number/
    var res1,res2 int
    for j:=0;j<len(nums);j++{
        if (nums[j]&div)==0{
            res1^=nums[j]
        }else{
            res2^=nums[j]
        }
    }
    return []int{res1,res2}
}
