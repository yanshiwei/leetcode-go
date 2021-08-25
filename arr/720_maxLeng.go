func longestWord(words []string) string {
    var fre=make(map[string]int)
    sort.Strings(words)
    var res string
    for i:=range words{
        word:=words[i]
        if len(word)==1{
            //single
            fre[word]++
            if len(res)==0{
                res=word
            }
        }
        //每次读取到len>1的字符串就到map里查询len-1的字符串是否存在
        if _,ok:=fre[word[0:len(word)-1]];ok{
            //word去掉末尾字符形成的单词在哈希表hashset中存在
            fre[word]++
            if len(res)<len(word){
                res=word
            }
        }
    }
    return res
}
