func maximalRectangle(matrix [][]byte) int {
    //以第i行为例，将第i行之前的对应列求和，最终得到所有列的高度数组如例子为例，第3行对应求和，则得到一个高度数组[3,1,3,2,2]如此转化成列84的高度数组的最大矩形问题故直接将其换成高度数组即可
    var res int
    if len(matrix)<1||len(matrix[0])<1 {
        return res
    }
    //get height
    var height=make([]int,len(matrix[0]))//对应列数即为数组长度
    for i:=0;i<len(matrix);i++ {
        for j:=0;j<len(matrix[0]);j++ {
            if matrix[i][j]=='1' {
                height[j]+=1
            }else {
                height[j]=0//新的一行为0，即使上一行该列非0也要重置
            }
        }
        //get height max rectangle for i
        res=max(res,getMaxRect(height))
    }
    return res
}

func getMaxRect(height[]int)int {
    var res int
    if len(height)<1 {
        return res
    }
    var stack []int
    height=append(height,0)//tail 0 handle res element,asc stack
    for i:=range height {
        for len(stack)>0&&height[i]<=height[stack[len(stack)-1]] {
            tmp:=height[stack[len(stack)-1]]
            stack=stack[0:len(stack)-1]//pop
            var area int
            if len(stack)>0 {
                area=tmp*(i-stack[len(stack)-1]-1)
            }else {
                area=tmp*i
            }
            res=max(res,area)
        }
        stack=append(stack,i)
    }
    return res
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
