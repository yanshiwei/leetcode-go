func maxChunksToSorted(arr []int) int {
    /*
    与769不同的是，存在重复元素，并且数组元素不再是 [1, n - 1] 的排列了，最大值可为 1e8。
    关于排序块的定义和条件：
    排序块定义：当某块后面的元素都大于此块内的元素，那么此块为排序块。每块长度最少为 1，即每个元素可单独成块。
    分成块的前提是：前一个块的最大值，小于后一个块的最小值

    要尽量分成多的块，使用贪心的思想，能分就分
    1.1 如果当前值 num 大于等于 前面块的最大值，那么 num 就可以单独成块
    1.2 如果当前值 num 小于前面块的最大值，那么我们需要往前一直找，直到找到一个块 block 的最大值小于等于当前值。而在这个 block 块 后面的块，都应该与 当前值 合并为一个块。因为它们的最大值都比 num 大，如果不合并为一个块，那么排序后肯定不是升序的状态。比如 num = 1，而前面一个块为 [2,4,1,6]，最大值为 6，如果 1 不和 前一块进行合并，那么排序后就变成 [1 2 4 6][1] = 1 2 4 6 1不是升序状态
    
    我们维护一个非严格的单调递增栈，栈内的元素是划分的每块中的最大值，即最终划分的块的个数为栈中元素个数。

    举例：
    arr = [1,1,2,1,1,3,4,1,6]

    遇到第一个 1，栈内为空，单独划分为一个块，入栈
    stack = [1]
    遇到第二个 1，等于 栈顶元素，单独划分为一个块，入栈
    stack = [1,1]
    遇到 2，大于 栈顶元素，单独划分为一个块，入栈
    stack = [1,1,2]
    遇到第一个 1，小于栈顶元素，意味着必须跟前面的块进行合并，我们将 2 弹出
    stack = [1,1]
    继续比较栈顶元素，等于栈顶元素，停止比较
    将弹出的块跟当前值进行合并，变成 [2,1] 块，最大值仍然是 2，将 2 压入栈
    stack = [1,1,2]
    遇到第二个 1，小于栈顶元素，跟上一步一样的做法
    跟前面的 [2,1] 块进行合并，变成 [2,1,1] 块
    stack = [1,1,2]
    遇到 3，大于栈顶元素，单独成块，入栈
    stack = [1,1,2,3]
    遇到 4，大于栈顶元素，单独成块，入栈
    stack = [1,1,2,3,4]
    遇到 1，小于栈顶元素，需要跟前面的块进行合并，我们弹栈弹栈弹栈
    直到 stack = [1,1]，当前值 等于 栈顶元素，因此停止弹栈
    跟 之前弹出的块进行合并，即 跟 [2,1,1] [3] [4] 三个块进行合并，变成 [2,1,1,3,4,1] 一个块
    最大值为 4，将 4 入栈
    stack = [1,1,4]
    遇到 6，大于栈顶元素，单独成块，入栈
    stack = [1,1,4,6]
    至此，最多可以划分为 4 个块
    */
    if len(arr)<1 {
        return 0
    }
    var stack []int//非严格递增栈
    for i:=range arr{
        tmp:=-1
        for len(stack)>0&&arr[i]<stack[len(stack)-1]{
            tmp=max(tmp,stack[len(stack)-1])
            stack = stack[0 : len(stack)-1]//pop
        }
        if tmp==-1 {
            stack=append(stack,arr[i])
        }else{
            stack=append(stack,tmp)
        }
    }
    return len(stack)
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
