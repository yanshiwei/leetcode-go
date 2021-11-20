var res []string
var path []string
func restoreIpAddresses(s string) []string {
    var length=len(s)
    res=nil
    if length<4||length>12{
        return res
    }
    path=nil//IP最多4段
    var splitTime int//已经得到切割多少段
    dfs(s,length,splitTime,0)
    return res
}

func dfs(s string, length int, splitTime int, begin int){
    // begin 待分割的起始位置
    if begin==length{
        if splitTime==4{
            //最多切割4次
            tmp:=append([]string(nil),path...)
            res=append(res,strings.Join(tmp,"."))
        }
        return
    }
     // 看到剩下的不够了，就退出（剪枝），len - begin 表示剩余的还未分割的字符串的位数. 
     // 剩下的IP 最少每个IP段区一位，最多取3位故4-splitTime～3*(4-splitTime)
     if length-begin<4-splitTime{
         return
     }
      // 看到剩下的超了，就退出（剪枝）
     if length-begin>3*(4-splitTime){
         return
     }
     for i:=0;i<3;i++{
         if begin+i>=length{
             // end=begin+i
             break
         }
         ipSeg,ok:=judgeIfIpSeg(s,begin,begin+i)
         if ok!=-1{
             // 在判断是 ip 段的情况下，才去做截取
             path=append(path,ipSeg)
             dfs(s,length,splitTime+1,begin+i+1)
             path=path[0:len(path)-1]
         }
     }
}

func judgeIfIpSeg(s string,begin int,end int)(string,int){
    var leng=end-begin+1
    // 大于 1 位的时候，不能以 0 开头
    if leng>1&&s[begin]=='0'{
        return "",-1
    }
    // 转成 int 类型
    str:=s[begin:end+1]
    i, err := strconv.Atoi(str)
    if err!=nil{
        return "",-1
    }
    if i>255{
        return "",-1
    }
    return str,0
}
