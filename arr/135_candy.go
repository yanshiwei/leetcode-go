func candy(ratings []int) int {
    /*
    相邻的孩子中，评分高的孩子必须获得更多的糖果拆解成：
    左规则：当 ratings[i−1]<ratings[i] 时，i号学生的糖果数量将比i−1 号孩子的糖果数量多
    右规则：当 ratings[i]>ratings[i+1] 时，i号学生的糖果数量将比i+1 号孩子的糖果数量多
    遍历该数组两次，处理出每一个学生分别满足左规则或右规则时，最少需要被分得的糖果数量。每个人最终分得的糖果数量即为这两个数量的最大值
    */
    var n=len(ratings)
    if n<1{
        return 0
    }
    var res int
    var left=make([]int,n)
    for i:=0;i<n;i++{
        if i>0&&ratings[i]>ratings[i-1]{
            left[i]=left[i-1]+1
        }else{
            left[i]=1
        }
    }
    var right int
    for i:=n-1;i>=0;i--{
        if i<n-1&&ratings[i]>ratings[i+1]{
            right++
        }else{
            right=1
        }
        res+=max(left[i],right)
    }
    return res
}
func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
