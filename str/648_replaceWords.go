func replaceWords(dictionary []string, sentence string) string {
    sort.Strings(dictionary)
    //保证最短优先替换
    sentenceList:=strings.Split(sentence," ")
    for i:=range sentenceList{
        one:=sentenceList[i]
        //每个单词可以在字典数组找到前缀时用字典中的替换
        for _,v:=range dictionary{
            if strings.HasPrefix(one,v){
                sentenceList[i]=v
                break
            }
        }
    }
    return strings.Join(sentenceList," ")
}
