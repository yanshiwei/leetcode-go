func isIdealPermutation(A []int) bool {
    /*
    局部倒置也是全局倒置；故只需要找到是全局倒置但非局部倒置即
    a[i]>a[j]&&j-i>1
    如果有 说明全局数多余局部
    */
    for i:=0;i<len(A);i++{
        for j:=i+2;j<len(A);j++{
            if A[i]>A[j]{
                return false
            }
        }
    }
    return true
}
