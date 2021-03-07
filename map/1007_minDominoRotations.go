func minDominoRotations(A []int, B []int) int {
    //定义一个3×6的数组，分别记录在某位置A数组单独出现数字1-6的次数，B数组单独出现数字1-6的次数，两数组出现相同数字1-6的次数.
    //之后在遍历6个数的出现次数，如果某个数i在三种情况出现的次数之和为n,n为A数组长，那么此时有解，只需旋转A和B单独出现i的次数较少的一个，更新答案
    var fre=make([][]int,3)
    for i:=range fre{
        fre[i]=make([]int,6)
    }
    for i:=range A {
        if A[i]==B[i]{
            fre[2][A[i]-1]++
        }else{
            fre[0][A[i]-1]++
            fre[1][B[i]-1]++
        }
    }
    var res =len(A)
    for i:=0;i<6;i++{
        if fre[0][i]+fre[1][i]+fre[2][i]==len(A){
            res=min(res,min(fre[0][i],fre[1][i]))
        }
    }
    if res==len(A){
        return -1
    }
    return res
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
