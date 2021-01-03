func leastInterval(tasks []byte, n int) int {
    /*
    1 先简单分析，假如只有一类任务如AAAA，n=2
    则处理流程：A 休闲 A 休闲 A 休闲
    除了最后一个A，其他除了都是任务处理时间1+间隔时间，最后一个A只需要处理时间不需要冷却故
    t=(任务数-1)*（1+n）+1
    2 对于有两种任务的，只要一种任务A的数量少于另一种B，那么A总可以穿擦在B的间隔中完成，故只需要看出现次数最多的任务的耗时
    3 根据2可以知道，不管有多少任务，只要小于最多的那个任务出现的次数，都总能穿插在间隔之间
    4 若存在任务和最多的任务相同，n = 2 输入为AAABBB,结果应该为
    A->B-> ->A->B-> ->A->B 结果是8
    除了最后A和B，其他的都是A和B交替，A和B交替执行，故只需要关注B的冷却时间故总耗时1+间隔时间，最后加A和B执行时间总和即可
    此时耗时t=(最多的那个任务的任务数 - 1)*（1+n）+ 任务数量等于最大的任务数的数量
    5 异常情况：
    AAABBBCCCDD n=2
    按照公式算出来结果是9，但实际上数组的长度都大于9，所以是特例，这种情况只需要返回数组长度就行了
    至于为什么会出现这样的情况，我们研究一下：
    他的一种可能结果是
    ABCDABCDABC,他不需要任何等待，就可以执行下一个任务，因此这种情况下需要返回数组长度
参考源头：
    链接：https://leetcode-cn.com/problems/task-scheduler/solution/         bu-yong-pai-xu-si-lu-jian-dan-qing-xi-su-rz8b/
    */
    var cnt=make([]int,26)//26个字母
    for i:=range tasks{
        cnt[tasks[i]-'A']++
    }
    var maxCnt=-1
    //find max
    for i:=0;i<26;i++{
        if maxCnt<cnt[i]{
            maxCnt=cnt[i]
        }
    }
    var res=(maxCnt-1)*(1+n)
    //考虑是否有多个
    for i:=0;i<26;i++{
        if maxCnt==cnt[i]{
            res++
        }
    }
    return max(res,len(tasks))//考虑异常
}
func max(x,y int)int{
    if x<y{
        return y
    }
    return x
}
