func longestConsecutive(nums []int) int {
    /*
    map
    每个数 x，考虑以其为起点，不断尝试匹配 x+1, x+2,⋯ 是否存在，假设最长匹配到了 x+y，那么以 x 为起点的最长连续序列即为 x, x+1, x+2, ,x+y，其长度为 y+1，我们不断枚举并更新答案.对于匹配的过程，暴力的方法是 O(n) 遍历数组去看是否存在这个数时间复杂度O(n^2)
    我们会发现其中执行了很多不必要的枚举，如果已知有一个 x, x+1, x+2, ⋯,x+y 的连续序列，而我们却重新从 x+1，x+2 或者是 x+y 处开始尝试匹配，那么得到的结果肯定比枚举 x 为起点的得到的长度要短，故碰到这种X存在前驱（即遍历X时存在X-1）情况时，直接不对X进行匹配跳过即可。
    用map存储数组中数字
    */
    var number=make(map[int]bool)
    for i:=range nums {
        if _,ok:=number[nums[i]];ok==false {
            number[nums[i]]=true
        }
    }
    var longest int
    for k,_:=range number{
        //判断有无前驱
        if _,ok:=number[k-1];ok==false {
            //如果没有前驱，则该值数一个连续序列的起点
            curNum:=k
            curLongest:=1
            ok1:=true
            for {
                ok1=number[curNum+1]
                if ok1==false {
                    break
                }
                curNum+=1
                curLongest+=1
            }
            longest=max(longest,curLongest)
        }
    }
    return longest
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
