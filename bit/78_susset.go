func subsets(nums []int) [][]int {
    /*
    一个包含n个元素的集合一共有2^n个子集
    先用位操作的思路来求解，具体方法：用2进制Bit位来标记集合中的某个元素是否被选中，1代表选中，0代表未选中。例如集合{a, b, c}的所有子集可如下表示：

    {}(空集)               0 0 0

    {a}                       0 0 1

    {b}                       0 1 0

    {c}                       1 0 0

    {a, b}                   0 1 1

    {a, c}                   1 0 1

    {b, c}                   1 1 0

    {a, b, c}               1 1 1
    */
    var m=len(nums)
    var mask int
    var start=0
    var end=(1<<m)-1
    var res[][]int
    for mask=start;mask<=end;mask++{
        one:=make([]int,0)
        for i:=0;i<m;i++ {
            if ((1<<i)&mask)!=0 {
                //1每次只出现在某一位，故每次选出一个
                one=append(one,nums[i])
            }
        }
        res=append(res,one)
    }
    return res
}
