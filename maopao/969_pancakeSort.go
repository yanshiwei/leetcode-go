func pancakeSort(arr []int) []int {
    /*
    maopao
    进行n次遍历，每次遍历找出当前区间[0,m]之间的最大数max及其下标j,
    每次遍历进行两次翻转，第一次将找到的最大数max翻转到最前面，第二次翻转最大数从最前面翻转到m处（即未排序位置的最后面）
    由于当前区间的最大数已经移动到合适位置（最后面),所以下一次遍历时的区间需要缩小1（n减去1
    */
    var m=len(arr)
    var res []int
    var cnt int
    for i:=0;i<m;i++{
        info:=getMax(arr[0:m-cnt])
        res=append(res,info.Index+1)
        //最大值反序到最前面
        reverse(arr[0:info.Index+1])
        res=append(res,m-cnt)
        //将最大值反序到未排序的最后
        reverse(arr[0:m-cnt])
        cnt++
    }
    return res
}

type Info struct {
    Data int
    Index int
}

func getMax(arr []int)Info{
    var res=Info{
        Data:INT_MIN,
        Index:-1,
    }
    for i:=range arr {
        if arr[i]>res.Data{
            res.Data=arr[i]
            res.Index=i
        }
    }
    return res
}

func reverse(arr []int){
    i,j:=0,len(arr)-1
    for i<j {
        arr[i],arr[j]=arr[j],arr[i]
        i++
        j--
    }
}

const INT_MAX=int(^(uint(0))>>1)
const INT_MIN=^INT_MAX
