func prefixesDivBy5(A []int) []bool {
    var res=make([]bool,len(A))
    var curV int
    for i:=0;i<len(A);i++{
        //考虑到数组 A可能很长，如果每次都保留curV的值，则可能导致溢出。由于只需要知道每个是否可以被 5 整除，因此在计算过程中只需要curV保留余数即可
        curV=((curV<<1)+A[i])%5
        if curV==0{
            res[i]=true
        }
    }
    return res
}
