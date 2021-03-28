func tupleSameProduct(nums []int) int {
    /*
    本例就是「两数之和」(https://leetcode-cn.com/problems/two-sum/)的拓展，查找使用哈希表,
    1.首先明确本题是((a, b), (c, d)) 组合个数，只要找到满足题意得a, b, c, d即可。其一共有 8 种组合！
        ((x), (y))：有两种组合，((x), (y))，((y), (x))
        (x) 组内调换顺序：(b, a), (a, b)
        (y) 组内调换顺序：(d, c), (c, d)
    因此总的组合有 2 * 2 * 2 共 8 种！（第一个2 表示两组(x)和(y)之间调换顺序，后面两个2 分别是(x),      (y)各组内调换顺序。
    2.查找这样的四元组对个数，使用「两数之和」思路即可，即一遍hash
    */
    var m=len(nums)
    var fre=make(map[int]int)
    var res int
    for i:=0;i<m-1;i++{
        for j:=i+1;j<m;j++{
            product:=nums[i]*nums[j]
            res+=8*fre[product]
            fre[product]+=1
        }
    }
    return res
}
