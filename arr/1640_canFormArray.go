func canFormArray(arr []int, pieces [][]int) bool {
    /*
    1.输入的数字范围是[0，100]，所以可以建立一个map[101]的数组，并存储它在arr的下标+1
    2.遍历pieces里面的数字，如果在map中没有记录，返回false，如果pieces中当前的行不只一个数，则判断相邻两个数是否在arr中从左往右连续，不符合则返回false
    3.排除所有false的情况后，则返回true
    */
    //1 <= pieces.length <= arr.length <= 100
    //1 <= arr[i], pieces[i][j] <= 100
    var fre=make([]int,101)
    for i:=0;i<len(arr);i++{
        fre[arr[i]]=i+1
    }
    for i:=0;i<len(pieces);i++{
        for j:=0;j<len(pieces[i]);j++{
            if fre[pieces[i][j]]==0{
                //如果在map中没有记录，返回false
                return false
            }
            if j>0{
                //判断相邻两个数是否在arr中从左往右连续
                delta:=fre[pieces[i][j]]-fre[pieces[i][j-1]]
                if delta!=1{
                    return false
                }
            }
        }
    }
    return true
}
