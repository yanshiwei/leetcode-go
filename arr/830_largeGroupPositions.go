func largeGroupPositions(s string) [][]int {
    var res [][]int
    if len(s)<3{
        return res
    }
    var m=len(s)
    var num int=1
    for i:=0;i<m;i++{
        if i==m-1{
            //到最后
            if num>=3{
                //push
                res=append(res,[]int{i-num+1,i})
            }
        }else if s[i]!=s[i+1]{
            //开始不同
            if num>=3{
                //push
                res=append(res,[]int{i-num+1,i})
            }
            num=1//reset
        }else{
            num++
        }
    }
    return res
}

