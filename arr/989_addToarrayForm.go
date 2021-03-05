func addToArrayForm(A []int, K int) []int {
    Karr:=intToArr(K)
    var mA=len(A)
    var mK=len(Karr)
    var m=min(mA,mK)
    var res=make([]int,max(mA,mK)+1)
    if mA<mK{
        assigin(Karr,res)
    }else{
        assigin(A,res)
    }
    var mRes=len(res)

    for i:=0;i<m;i++{
        if mA<mK{
            res[mRes-1-i]+=+A[mA-1-i]
        }else {
            res[mRes-1-i]+=+Karr[mK-1-i]
        }
        if res[mRes-1-i]>=10{
            res[mRes-1-i]-=10
            res[mRes-1-i-1]+=1
        }
    }
    for i:=len(res)-1;i>0;i--{
        if res[i]>=10{
            res[i]-=10
            res[i-1]+=1
        }
    }
    if res[0]==0{
        res=res[1:]
    }
    return res
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}

func assigin(src,dest []int){
    for i:=0;i<len(src);i++{
        dest[i+1]=src[i]
    }
}

func intToArr(a int)[]int {
    var res[]int
    aStr:=strconv.Itoa(a)
    for i:=range aStr{
        res=append(res,int(aStr[i]-'0'))
    }
    return res
}
