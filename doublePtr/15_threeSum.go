func threeSum(nums []int) [][]int {
	/*
	   采用排序+双指针的思路：
	   0 长度低于3 直接返回
	   1 首先先将给定的数组 nums 排序
	   2 遍历数组，固定三个数中第一个
	       2.1 若nums[i]>0 因为已经从小到大排序，加上后面的肯定也是大于0，故不用看后面的，直接返回
	       2.2 对nums[i]这个第一个数去重
	       2.3 双指针L，R在nums[i]后面的区间中寻找和为0-nums[i]的另外两个数
	           2.3.1 若后两个数之和大于nums[i]，说明两数之和太大，右指针左移
	           2.3.2 若后两个数之和小于nums[i]，说明两数之和太小，左指针右移
	           2.3.3 若后两个数之和等于nums[i]，找到一个解，并同时将L,R移到下一位置，寻找新的解；并对后两个去重，判断左指针和右指针是否和上一个位置重复
	*/
	var res[][]int
	if len(nums)<3 {
		return res
	}
	sort.Ints(nums)
	//从小到大
	for i:=range nums{
		//固定第一个数
		if nums[i]>0{
			return res
		}
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		left:=i+1
		right:=len(nums)-1
		for left<right{
			if nums[left]+nums[right]>0-nums[i]{
				//说明两数之和太大，右指针左移
				right--
			}else if nums[left]+nums[right]<0-nums[i]{
				//说明两数之和太小，左指针右移
				left++
			}else {
				res=append(res,[]int{nums[i],nums[left],nums[right]})
				left++
				right--
				//去重
				for left<right&&nums[left]==nums[left-1]{left++}
				for left<right&&nums[right]==nums[right+1]{right--}
			}
		}
	}
	return res
}
