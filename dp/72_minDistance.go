func minDistance(word1 string, word2 string) int {
    //dp[i][j] 代表 word1 中前 i 个字符，变换到 word2 中前 j 个字符，最短需要操作的次数
    //需要考虑 word1 或 word2 一个字母都没有，即全增加/删除的情况，所以预留 dp[0][j] 和 dp[i][0]
    var m1=len(word1)
    var m2=len(word2)
    var dp=make([][]int,m1+1)
    for i:=range dp{
        dp[i]=make([]int,m2+1)
    }
    for i:=0;i<len(dp);i++{
        //word 2 is empty, word1 delete one by one
        dp[i][0]=i
    }
    for j:=0;j<len(dp[0]);j++{
        // word 1 is empty, word1 add one by one
        dp[0][j]=j
    }
    for i:=1;i<len(dp);i++{
        for j:=1;j<len(dp[0]);j++{
            if word1[i-1]==word2[j-1]{
                //pos i-1 is the same as pos j-1, thus we focus on one before
                dp[i][j]=dp[i-1][j-1]
            }else{
                //add,word1 pos i-1 add one element
                add:=dp[i][j-1]+1
                //delete,word1 pos i-1 delete one element
                del:=dp[i-1][j]+1
                //update,word1 pos i-1 update one element
                upd:=dp[i-1][j-1]+1
                dp[i][j]=min(add,min(del,upd))
            }
        }
    }
    return dp[m1][m2]
}

func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
