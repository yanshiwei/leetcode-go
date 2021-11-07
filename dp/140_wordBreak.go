func wordBreak(s string, wordDict []string) []string {
    //https://leetcode-cn.com/problems/word-break-ii/solution/dong-tai-gui-hua-jie-jue-qing-xi-yi-dong-ndrf/
    var slen=len(s)
    var dp=make([][]string,slen+1)//第i个字符时前面所能拆分的句子列表
    dp[0]=[]string{""} //第0个字符肯定是空结果
    var wordFre=make(map[string]struct{})
    for i:=0;i<len(wordDict);i++{
        wordFre[wordDict[i]]=struct{}{}
    }
    for i:=1;i<=slen;i++{
        // loop for exist
        for j:=i-1;j>=0;j--{
            if len(dp[j])<1{
                continue
            }
            suffix:=s[j:i]
            if _,ok:=wordFre[suffix];ok{
                if len(dp[j])>0{
                    // fmt.Println("s",i,j,suffix,(dp[j]))
                    for k:=0;k<len(dp[j]);k++{
                        if j==0{
                            tmp:=dp[j][k]+""+suffix
                            dp[i]=append(dp[i],tmp)
                        }else{
                            // 非第一个字符 则空格隔开
                            tmp:=dp[j][k]+" "+suffix
                            dp[i]=append(dp[i],tmp)
                        }
                    }
                }
            }
        }
    }
    return dp[slen]
}
