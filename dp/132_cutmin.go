func minCut(s string) int {
    //https://leetcode-cn.com/problems/palindrome-partitioning-ii/solution/fen-ge-hui-wen-chuan-ii-by-leetcode-solu-norx/
    var n=len(s)
    //O(1) 的时间就可以判断任意 s[i..j]是否为回文串
    //dp 将字符串 s 的每个子串是否为回文串预先计算出来
    var g=make([][]bool,n)
    for i:=range g{
        g[i]=make([]bool,n)
        for j:=range g[i]{
            g[i][j]=true
        }
    }
    //转移方程里需要g[i+1][j-1]，这个时候i+1还没判断，所以要把i从n-1减小到0
    for i:=n-1;i>=0;i--{
        fmt.Println(i,i+1)
        for j:=i+1;j<n;j++{
            
            if s[i]==s[j]{
                g[i][j]=g[i+1][j-1]
            }else{
                g[i][j]=false
            }
        }
    }
    var f=make([]int,n)//字符串的前缀 s[0..i]的最少分割次数
    for i:=range f{
        if g[0][i]{
            // s[0...i]已经是回文
            f[i]=0
            continue
        }
        f[i]=math.MaxInt64
        for j:=0;j<i;j++{
            if g[j+1][i]&&f[j]+1<f[i]{
                // s[j+1:...,i]是回文
                f[i]=f[j]+1
            }
        }
    }
    return f[n-1]
}
