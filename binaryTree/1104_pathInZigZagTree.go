func pathInZigZagTree(label int) []int {
    /*每一层值x取值范围2^(level-1)<=x<2^(level),其中level是高度从1开始
    对于奇数层，从2^(level-1)开始，从小到大，故下标=x-2^(level-1)
    对于偶数层，从2^(level)-1开始，从大到小，故下标=2^(level)-1-x
    下标从0开始计算, n-1层下标 = n层下标/2
    1.可以先假设全部从左到右升序，那么很容易看出 父节点 = 本节点 / 2，如例子中14的父节点为7
    2.由于奇偶层逆序，原来的父节点就变成轴对称的另一边去了 7变成4 6变成5；但是基于轴对称的和仍是固定的如15+8=23；14+9=24...，故可以假设新增一行，其中一行从小到大，一行从大到小，如例子：
    已有行15 14 13 12 11 10  9  8
    新增行8  9  10 11 12 13 14  15
    如只考虑新增行，则生序父节点=本节点 / 2，但是奇偶逆序，此时真实的父节点=固定值-生序父节点
    固定值可以由该层最值和计算
    故父节点 = (上层最大值 + 上层最小值) - 本节点 / 2
    */
    var level=0
    //判断label所在行
    for (1<<level)-1<label{
        level++
    }
    //fmt.Println(level)
    var res []int
    res=append(res,label)
    for label!=1{
        level--
        // 以14为例，父节点在第3层
        // 上层最大值 = 2 ^ 3 - 1 = 7，上层最小值 = 2 ^ 2 = 4
        // 父节点 = 7 + 4 - 14/2
        maxV:=(1<<level)-1
        minV:=(1<<(level-1))
        label=maxV+minV-label/2
        res=append(res,label)
    }
    var ans []int
    for i:=len(res)-1;i>=0;i--{
        ans=append(ans,res[i])
    }
    return ans
}
