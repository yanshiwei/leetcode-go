var res [][]int
func permute(nums []int) [][]int {
    res=nil
    helper(nums,[]int(nil))
    return res
}

func helper(nums[]int,tmp []int){
    if len(tmp)==len(nums){
        //不能直接用tmp
        // 这个 tmp 变量是一个地址引用，结束当前递归，将它加入 res，后续的递归分支还要继续进行搜		
		// 索，还要继续传递这个 tmp ，这个地址引用所指向的内存空间还要继续被操作，所以 res 中的 		 
		// tmp 所引用的内容会被改变，这就造成了 res 中的内容随 tmp 变化。
        // 所以要复制 tmp 内容到一个新的数组里，然后放入 res，这样后续对 tmp 的操作，
		// 就不会影响已经放入 res 的内容。 
        temp := make([]int, len(tmp))
        copy(temp,tmp)
        res=append(res,temp)
        return
    }
    for i:=0;i<len(nums);i++{
        if container(tmp,nums[i]){
            continue
        }
        tmp=append(tmp,nums[i])
        //next
        helper(nums,tmp)
        //recover
        tmp=tmp[:len(tmp)-1]
    }
}

func container(nums[]int,num int)bool{
    for i:=0;i<len(nums);i++{
        if nums[i]==num{
            return true
        }
    }
    return false
}
