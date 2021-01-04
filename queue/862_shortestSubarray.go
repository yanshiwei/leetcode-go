func shortestSubarray(A []int, K int) int {
    /*
     原数组 ：【1, 2, -2, 3, 1,-4, 2】,转化为前缀和数组，即当前位置之前所有元素的和
    转化后 ：【1, 3, 1, 4, 5 , 1, 3】

    问题转化为：前缀和数组两个元素的差值为K，距离的最小间隔。
    类似单调栈，通过单调增队列解决，单调队列队尾进元素时和单调栈一样，对于破坏单调性的元素直接弹出队尾
    */
    var  queueIdx =make([]int,50001)
    var Alen=len(A)
    var head =0//队列首部位置
    var tail=0//队列尾部的下一个位置
    if Alen<1{
        return -1
    }
    var sumPrefixA=make([]int,Alen+1)
    sumPrefixA[0]=0
    for i:=0;i<Alen;i++{
        sumPrefixA[i+1]=sumPrefixA[i]+A[i]
    }
    var res=50001
    for i:=0;i<Alen+1;i++{
        for head<tail{
            //队列非空
            if sumPrefixA[i]<sumPrefixA[queueIdx[tail-1]]{
                //最新元素比队尾对应的元素小，影响递增单调性，出队列，直到满足单增位置
                tail--
            }else{
                break
            }
        }
        queueIdx[tail]=i//满足递增性的最新元素的index压入队列中
        tail++
        //如此对内对应元素都是单增的
        for head<tail{
            //队列非空  从头部 差值可以最大 反复寻找差值大于K的  直到不满足为止 最后得到的j-i差值最小
            if sumPrefixA[i]-sumPrefixA[queueIdx[head]]>=K{
                res=min(res,i-queueIdx[head])//贪心
                head++
            }else{
                break
            }
        }
    }
    if res<=Alen{
        return res
    }
    return -1
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
