func removeSubfolders(folder []string) []string {
    var m=len(folder)
    var root=NewTriNode()
    sort.Slice(folder,func(i,j int)bool{
        x:=folder[i]
        y:=folder[j]
        n:=len(x)
        m:=len(y)
        //按照字典序
        var idx int
        for idx=0;idx<n&&idx<m;idx++{
            if x[idx]<y[idx] {
                return true
            }else if x[idx]>y[idx]{
                return false
            }
        }
        //长度一样
        if idx==n&&idx==m {
            return true
        }
        //谁长谁大
        if idx<n||idx<m{
            return n<m
        }
        return true
    })
    var res []string
    for i:=0;i<m;i++{
        tmp:=folder[i]+"/"
        if root.TryInsert(tmp)==false{
            res=append(res,folder[i])
        }
    }
    return res
}

type TriNode struct {
    Path int//有多少单词公用该节点
    IsWord bool//有多少单词以这个节点结尾
    Next []*TriNode
}

func NewTriNode()*TriNode{
    return &TriNode{
        Next:make([]*TriNode,27),
        //include '/'
    }
}

func (this *TriNode)Insert(word string){
    if len(word)==0{
        return
    }
    cur:=this
    for i:=range word{
        if cur.Next[int(word[i]-'a')]==nil{
            //new
            cur.Next[int(word[i]-'a')]=NewTriNode()
        }
        cur=cur.Next[int(word[i]-'a')]
    }
    if cur.IsWord==false{
       cur.IsWord=true
       cur.Path++ 
    }
}

func (this *TriNode)TryInsert(word string)bool{
    if len(word)==0{
        return true
    }
    cur:=this
    for i:=range word{
        var pos =26
        if word[i]!='/'{
            pos=int(word[i]-'a')
        }
        if cur.Next[pos]==nil{
            //new
            cur.Next[pos]=NewTriNode()
        }
        cur=cur.Next[pos]
        if cur.IsWord {
            return true
        }
    }
    if cur.IsWord==false{
       cur.IsWord=true
       cur.Path++ 
    }
    return false
}

//是否包含某单词
func (this *TriNode)Find(word string)bool {
    if len(word)==0{
        return false
    }
    cur:=this
    for i:=range word{
        if cur.Next[int(word[i]-'a')]==nil{
            return false
        }
    }
    return cur.IsWord
}

//是否包含preWith前缀的单词
func (this *TriNode)FindPre(preWith string)bool{
    if len(preWith)==0{
        return false
    }
    cur:=this
    for i:=range preWith{
        if cur.Next[int(preWith[i]-'a')]==nil{
            return false
        }
    }
    return true
}

