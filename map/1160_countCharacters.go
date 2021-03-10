func countCharacters(words []string, chars string) int {
    /*
    对于一个单词 word，只要其中的每个字母的数量都不大于 chars 中对应的字母的数量，那么就可以用 chars 中的字母拼写出 word
    */
    var charsFre=make(map[byte]int)
    for i:=range chars{
        charsFre[chars[i]]++
    }
    var res int
    for i:=range words{
        word:=words[i]
        isFind:=true
        var wordFre=make(map[byte]int)
        for j:=range word {
            wordFre[word[j]]++
        }
        //fmt.Println(word,charsFre,wordFre)
        for j:=range word{
            if wordFre[word[j]]>charsFre[word[j]]{
                isFind=false
                break
            }
        }
        if isFind{
            //fmt.Println(word)
            res+=len(word)
        }
    }
    return res
}
