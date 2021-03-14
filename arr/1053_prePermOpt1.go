func prevPermOpt1(arr []int) []int {
    //从后向前寻找第一个非严格递减的数x，然后在寻找最大的小于x的数，两者交换
    if len(arr)<1 {
        return arr
    }
    var firstPos=-1
    var lastV=arr[len(arr)-1]
    for i:=len(arr)-1;i>=0;i--{
        if arr[i]>lastV{
            firstPos=i
            break
        }
        lastV=arr[i]
    }
    if firstPos==-1 {
        //not found 
        return arr
    }
    var firstLessPos=-1
    //从firstPos开始寻找比它小的第一个最大数
    for i:=len(arr)-1;i>firstPos;i--{
        if arr[i]<arr[firstPos]{
            firstLessPos=i
            break
        }
    }
    lastV=arr[firstLessPos]
    for arr[firstLessPos]==lastV{
        firstLessPos--
    }
    firstLessPos++
    arr[firstPos],arr[firstLessPos]=arr[firstLessPos],arr[firstPos]
    return arr
}

const INT_MAX=int(^uint(0)>>1)
const INT_MIN=^INT_MAX
