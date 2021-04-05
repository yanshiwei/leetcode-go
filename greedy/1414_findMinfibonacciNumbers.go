func findMinFibonacciNumbers(k int) int {
    //贪心，根据k先构造一个计算到大于k的斐波那契项，用数组存着所有项的值，肯定是递增的，然后从数组尾部开始减，我们每次都用最大数（最好是等于k的数嘛）去减k，肯定总共能减的次数就会最少。题目已经说明：数据保证对于给定的 k ，一定能找到可行解
    if k<3{
        return 1
    }
    fb := []int{1,1}
    for i := 1; fb[i] < k; i++ {
        fb = append(fb, fb[i] + fb[i-1])
    }
    //fb 一个单调递增序列
    //fmt.Println(fb)
    var res int
    for i:=len(fb)-1;i>=0;i--{
        if fb[i]>k {
            continue
        }
        k-=fb[i]
        res++
        if k==0{
            return res
        }
    }
    return res
}
//不要通过函数调用，会存在开销
func fib(n int)int {
    if n<3{
        return 1
    }
    return fib(n-1)+fib(n-2)
}
