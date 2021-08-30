var res [][]int
var used[]bool
/*
同一轮中即使前后数字相同（即重复）也可以加，但是不同轮中不行。以[1,1,2]为例 在第一轮（初始轮）中，先取了第一个1之后，第二个1依然可以取，并不能被continue掉，说明单纯 if (vis[i] || (i > 0 && nums[i] == nums[i - 1])) { continue; }是不行的，因为这样的话就会执行continue了。 而什么时候需要continue呢？实际上是第一轮结束后，[1,1,2]已经被加入结果集中之后，最外层for循环（也就是取可行解第一个元素的for循环）往后运行，走到下标为1的第二个“1”的位置，我称之为第二轮，在第二轮中先确定可行解第一个元素为“1”之后再调用backtrack进入第二层for循环，for循环依然从i=0开始走，会发现此时0号位的1是false，那么这个“1”加不加呢？答案是不能加，因为假设加了之后，第二轮的可行解的[1,1,2]跟第一轮的可行解[1,1,2]是重复的（唯一的区别就是这里面的1和1对应原数组的下标是互换的）。 所以，!vis[i - 1]这个限制条件的作用主要是针对不同轮。

所以，如果两个相同的数字中（为了容易区分，命名为1a 和 1b），前一个（1a）的标记位为false，说明此时肯定不是在同一轮了（因为每时每刻处于同一轮的一定是标记位同时为true的（好好体会这句话），况且1a还在前面，如果在同一轮中，1a肯定是true，之所以为false，说明是回溯到1a的父层后， 先把1a置false，再for遍历到1b），既然不在同一层，那就要注意，不能在1b的轮中再加入1a了（因为1a的轮中肯定加入过1b了），所以判断1a和1b值相同，同时1a标记位为false，那就果断continue，那如果1b后面还有1c呢？同样的道理，只需要判断1b的状态是不是false，如果是false，说明他俩不在同一轮，1b的轮次中肯定加入过1c了，所以1c的轮次中就可以跳过1b
*/
func permuteUnique(nums []int) [][]int {
    res=nil
    used=make([]bool,len(nums))
    sort.Ints(nums)
    helper(nums,[]int(nil))
    return res
}
func helper(nums[]int,tmp[]int){
    if len(nums)==len(tmp){
        temp:=make([]int,len(tmp))
        copy(temp,tmp)
        res=append(res,temp)
        return
    }
    for i:=0;i<len(nums);i++{
        if used[i]{
            //我们已经选择过的不需要再放进去
            continue
        }
        if i>0&&nums[i]==nums[i-1]&&used[i-1]==false{
            //每次填入的数一定是这个数所在重复数集合中「从左往右第一个未被填过的数字」
            continue
        }
        tmp=append(tmp,nums[i])
        used[i]=true
        helper(nums,tmp)
        //recove
        tmp=tmp[:len(tmp)-1]
        used[i]=false
    }
}
