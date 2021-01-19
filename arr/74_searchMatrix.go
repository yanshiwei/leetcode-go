func searchMatrix(matrix [][]int, target int) bool {
    //看作一个长m*n的有序数组，数组下标idx
    //则row=idx/n col=idx%n
    var m=len(matrix)
    var n=len(matrix[0])
    var left=0
    var right=m*n-1
    var mid int
    for left<=right{
        mid=left+(right-left)/2
        element:=matrix[mid/n][mid%n]
        if target==element {
            return true
        }else {
            if target<element {
                right=mid-1
            }else {
                left=mid+1
            }
        }
    }
    return false
}
