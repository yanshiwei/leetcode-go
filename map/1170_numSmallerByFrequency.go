func numSmallerByFrequency(queries []string, words []string) []int {
    //题意是queries里面每个单词从字典words中查找符合要求f(q) < f(w)的数量
    //注意words可能存在重复元素
    var res []int
    var maxWordFre=make(map[string]Info)
    for i:=range words{
        word:=words[i]
        wordFre:=getMinFre(word)
        var info Info
        if _,ok:=maxWordFre[word];ok==false{
            info=Info{Fre:wordFre,Cnt:1}
        }else{
            info=maxWordFre[word]
            info.Cnt++
        }
        maxWordFre[word]=info
    }
    //fmt.Println(len(words),maxWordFre)
    for i:=range queries{
        query:=queries[i]
        funOfquery:=getMinFre(query)
        var cnt int
        for _,v:=range maxWordFre{
            if funOfquery<v.Fre{
                cnt+=v.Cnt
            }
        }
        res=append(res,cnt)
    }
    return res
}

func getMinFre(s string)int{
    var minByte=s[0]
    for i:=range s{
        if minByte>s[i]{
            minByte=s[i]
        }
    }
    var minByteFre int
    for i:=range s{
        if s[i]==minByte{
            minByteFre++
        }
    }
    return minByteFre
}

type Info struct {
    Fre int
    Cnt int
}
