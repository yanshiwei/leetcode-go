func numTeams(rating []int) int {
    var n=len(rating)
    var res int
    /*
    超时
    for i:=0;i<n;i++{
        for j:=i+1;j<n;j++{
            for k:=j+1;k<n;k++{
                if rating[i]<rating[j]&&rating[j]<rating[k]{
                    res++
                }else if rating[i]>rating[j]&&rating[j]>rating[k]{
                    res++
                }
            }
        }
    }
    */
    /*
    0.某个士兵j为中间的士兵,存在两种情况：左侧比他小，右侧比他大；左侧比他大，右侧比他小；
    1.假设以某个士兵j为中间的士兵，分别统计其 左/右 侧评分 大于/小于 rating[j]的士兵数，则以士兵j为中间的士兵可以组成的作战单位数量为smaller_cnt_l*bigger_cnt_r + bigger_cnt_l*smaller_cnt_r;
    2.遍历所有的士兵，以该士兵为中间，统计并加和其可以组成的作战单位数量，整体遍历一遍后，所有的组合结果肯定无重复，因此这就是答案
    */
    for j:=1;j<n-1;j++{
        var smaller_cnt_l int//左侧比该值小
        var bigger_cnt_l int//左侧比该值大
        var smaller_cnt_r int//右侧比该值小
        var bigger_cnt_r int//右侧比该值大
        //左侧
        for i:=0;i<j;i++{
            if rating[i]<rating[j]{
                smaller_cnt_l++
            }else{
                bigger_cnt_l++
            }
        }
        //右侧
        for k:=j+1;k<n;k++{
            if rating[k]<rating[j]{
                smaller_cnt_r++
            }else{
                bigger_cnt_r++
            }
        }
        //针对该士兵j，满足条件的有smaller_cnt_l*bigger_cnt_r+bigger_cnt_l*smaller_cnt_r
        res+=smaller_cnt_l*bigger_cnt_r+bigger_cnt_l*smaller_cnt_r
    }
    return res
}
