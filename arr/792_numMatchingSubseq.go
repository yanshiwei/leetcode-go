func numMatchingSubseq(S string, words []string) int {
    var cnt int
    fmt.Println(len(words))
    for i:=range words{
        notFind:=false
        pos:=-1
        word:=words[i]
        for j:=range word{
            var letter string
            letter=word[j:j+1]
            pos=findFirstPos(S,letter,pos+1)
            if pos==-1{
                notFind=true
                break
            }else{
                notFind=false
            }
        }
        if notFind==false{
            //表示word已全部被遍历了，则为其子串
            cnt++
        }
    }
    return cnt
}
func findFirstPos(s,sub string,pos int)int{
    finalPos:=strings.Index(s[pos:],sub)
    if pos==-1{
        return finalPos
    }else{
        if finalPos!=-1{
            return finalPos+pos
        }
    }
    return finalPos
}

