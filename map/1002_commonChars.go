func commonChars(A []string) []string {
    var res []string
    var fre=make([]int,26)//每个字符在某个单词出现的次数
    var minFre=make([]int,26)//每个字符遍历所有单词出现的最少次数
    for i:=range minFre{
        minFre[i]=INT_MAX
    }
    for i:=range A {
        for i:=range fre{
            fre[i]=0
        }//清空
        word:=A[i]
        for j:=0;j<len(word);j++{
            fre[word[j]-'a']++
        }
        for j:=0;j<26;j++{
            minFre[j]=min(minFre[j],fre[j])
        }
        //获取当前每个字符的最小次数
    }
    for j:=0;j<26;j++{
        for k:=0;k<minFre[j];k++{
            res=append(res,string(j+'a'))
        }
    }
    return res
}
const INT_MAX=int(^(uint(0))>>1)

func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
