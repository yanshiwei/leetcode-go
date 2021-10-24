func isScramble(s1 string, s2 string) bool {
    var length int=len(s1)
    if length!=len(s2){
        return false
    }
    var dp=make([][][]bool,length)
    for i:=range dp{
        dp[i]=make([][]bool,length)
        for j:=range dp[i]{
            dp[i][j]=make([]bool,length+1)
        }
    }
    //dp[k][i][j]表示s1从第i个字符开始，s2从第j个字符开始，他们分别截取长度为k的子字符串，判断这两个子字符串是否相互为扰乱字符串
    for k:=1;k<=length;k++{
        var end=length-k
        for i:=0;i<=end;i++{
            for j:=0;j<=end;j++{
                if k==1{
                    //如果比较的长度是1，也就是一个字符，直接判断他们是否相等即可
                    dp[i][j][k]=s1[i]==s2[j]
                    continue
                }else{
                    //这k个字符不断的截成两部分，然后s1的前面部分和s2的前面部分，s1的后面部分和s2的后面部分在计算。或者s1的前面部分和s2的后面部分，s1的后面部分和s2的前面部分计算。
                    //随机，故left size 可以from 1 to k-1 to try till has matched
                    for m:=1;m<k;m++{
                        var leftToLeft=dp[i][j][m]&&dp[i+m][j+m][k-m]//s1左边--->s2左边，长度m      s1右边--->s2右边，长度k-m
                        var leftToTight=dp[i][j+k-m][m]&&dp[i+m][j][k-m]//s1左边--->s2右边，长度m      s1右边--->s2左边，长度k-m
                        //扰乱字符串，上面只要有一个匹配即可
                        dp[i][j][k]=leftToLeft||leftToTight
                        if dp[i][j][k]{
                            break
                        }
                    }
                }
            }
        }
    }
    return dp[0][0][length]
}
