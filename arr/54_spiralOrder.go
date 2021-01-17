func spiralOrder(matrix [][]int) []int {
    /*
    1.按照顺序遍历上右下左四条边，不停缩小四条边范围
    2.定义一个计数变量idx,当idx等于所有元素数量之和则跳出循环
    3.定义表示四条边的变量left,right,top,bottom表示左右上下边界
    4.每遍历一条边该范围缩小1，比如上边遍历完，则t减一，左边遍历完则l加1重复这个过程
注意有可能在循环结束前已经超过了范围，所以在遍历四条边时候，也要加上条件num <= toal
    */
    var left=0
    var right=len(matrix[0])-1
    var top=0
    var bottom=len(matrix)-1
    var totalNum=len(matrix)*len(matrix[0])
    var idx int
    var cache []int
    for idx<totalNum{
        //从左到右
        for i:=left;i<=right&&idx<totalNum;i++{
            cache=append(cache,matrix[top][i])
            idx++
        }
        top++
        //从上到下
        for i:=top;i<=bottom&&idx<totalNum;i++{
            cache=append(cache,matrix[i][right])
            idx++
        }
        right--
        //从右到左
        for i:=right;i>=left&&idx<totalNum;i--{
            cache=append(cache,matrix[bottom][i])
            idx++
        }
        bottom=bottom-1
        //从下往上
        for i:=bottom;i>=top&&idx<totalNum;i--{
            cache=append(cache,matrix[i][left])
            idx++
        }
        left++
    }
    return cache
}
