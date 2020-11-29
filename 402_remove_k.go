func removeKdigits(num string, k int) string {
    //从左到右找到第一个递减的位置，将前面的大数删掉
    //如果已经是非递减序列，直接去掉最后K个
    var res="0"
    if len(num)<1 {
        return res
    }
    if len(num)==1&&k==1 {
        return res
    }

    var st []int//st is not desc
    var deletNum int
    var i int
    for i=range num {
        v:=int(num[i]-'0')
        if len(st)<1 {
            st=append(st,v)
        }else {
            if v<st[len(st)-1] {
                //第一个递减的
                for v<st[len(st)-1] && deletNum<k {
                    st=st[0:len(st)-1]//pop
                    deletNum++
                    if len(st)<1 {
                        break
                    }
                }
                st=append(st,v)
            }else {
                st=append(st,v)
            }
        }
    }

    //empty
    if len(st)<1 {
        return res
    }

    if deletNum<k {
        for i=0;i<k-deletNum;i++ {
            st=st[0:len(st)-1]//pop
        }
    }

    //delete pre 0
    i=0
    for len(st)>0&&st[0]==0 {
        st=st[1:]//pop head
    }
    res=""
    for i=range st {
        str:=strconv.Itoa(st[i])
        res+=str
    }
    return res
}
