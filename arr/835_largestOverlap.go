func largestOverlap(img1 [][]int, img2 [][]int) int {
    /*
    A中的1所在的点去移动到B中1所在的点。之后去判断A中所有的1所在的位置加上这段位移之后是否与B之中的1位置重合。
    举例来说，A中第一个为1的点A1（0，0），去找B中第一个1点（1，1），需要的位移是右移1，下移1，那么我们让A中所有值为1的点都加上这个位移，然后每个去判断是否与B中值为1点的位置重合。
    */
    var m=len(img1)
    var A,B []Node
    for i:=0;i<m;i++{
        for j:=0;j<m;j++{
            if img1[i][j]==1 {
                A=append(A,Node{i,j})
            }
            if img2[i][j]==1{
                B=append(B,Node{i,j})
            }
        }
    }
    var Bset=make(map[Node]struct{})
    for i:=range B {
        Bset[B[i]]=struct{}{}
    }
    var seen=make(map[Node]struct{})
    var res int
    for i:=range A{
        //A中每个1去与B的每个1去重合
        for j:=range B {
            //这个delta可以理解为A中为1的点a要走到B中为1的b这个点需要走多少。
            //like A中第一个为1的点A1（0，0），去找B中第一个1点（1，1），需要的位移是右移1，下移1
            delta:=Node{B[j].X-A[i].X,B[j].Y-A[i].Y}
            if _,ok:=seen[delta];ok==false{
                //为了避免相同的位移。比如A中（0，1）处的1想到B中（1，2）处的1也是需要向右移动1，向下移动1
                seen[delta]=struct{}{}
                //对于A中每个为1的点加上位移，判断是否与B为1的重叠
                var oneDeltaNum int
                for p:=range A{
                    newPos:=Node{A[p].X+delta.X,A[p].Y+delta.Y}
                    if _,ok1:=Bset[newPos];ok1==true{
                        oneDeltaNum++
                    }
                }
                res=max(res,oneDeltaNum)
            }
        }
    }
    return res
}

func max(x,y int)int{
    if x<y {
        return y
    }
    return x
}

type Node struct {
    X int
    Y int
}
