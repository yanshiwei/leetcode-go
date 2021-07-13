var C [][]int
func numOfWays(nums []int) int {
    var n=len(nums)
    //组合数组
    C=make([][]int,n+1)
    for i:=range C {
        C[i]=make([]int,n+1)
    }
    //初始化组合数组
    for i:=0;i<=n;i++{
        for j:=0;j<=i;j++{
            //n>=m
            if j==0{
                C[i][j]=1
            }else{
                C[i][j]=(C[i-1][j-1]+C[i-1][j])%mod
            }
        }
    }
    // 排除题中给出的一种
    return (f(nums) + mod - 1) % mod
}

func f(nums []int)int {
    if len(nums)==0{
        return 1
    }
    var n=len(nums)
    var rootV=nums[0]
    var left []int
    var right []int
    for i:=range nums{
        if nums[i]>rootV{
            right=append(right,nums[i])
        }
        if nums[i]<rootV{
            left=append(left,nums[i])
        }
    }
    return C[n-1][len(left)]*f(left)%mod*f(right)%mod
}

const mod int =1e9+7
/*
1.f(nums)的逻辑: 第一个值作为根是固定的，小于根在左(设共有x个)，大于在右(nums.size()-1-x个)，根的两边的数互相不会干扰，体现在数组上则是，nums[0]为根,nums[1:最后]中选择x个位置放小于根的数(共C[nums.size()-1][x]种)
2.子问题f(smaller_nums)：小于或者大于根的子树，则成为了原问题的递归问题(原问题为f(all_nums), 子问题则为f(smaller_nums))
3.结合1与2，最终结果则是 f(all_nums) = C[all_nums.size()-1][small_nums.size()] * f(small_nums) * f(bigger_nums);
4.C(n,m)=n!/(m!(n-m)!)=C(n-1,m-1)+C(n-1,m),n>=m
*/
