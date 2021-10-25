func groupAnagrams(strs []string) [][]string {
    //字母异位词指字母相同，但排列不同的字符串
    var fre=make(map[string][]string)
    for i:=range strs{
        s:=[]byte(strs[i])
        sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
        sortedStr:=string(s)
        fre[sortedStr]=append(fre[sortedStr],strs[i])
    }
    var res [][]string
    for _,v:=range fre{
        res=append(res,v)
    }
    return res
}
